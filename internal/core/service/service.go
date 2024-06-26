package service

import (
	"fmt"
	"github.com/PavlushaSource/comp-club-task/internal/core/entity"
	"github.com/PavlushaSource/comp-club-task/internal/core/utils"
	"github.com/PavlushaSource/comp-club-task/internal/lib/parser"
	"log"
	"slices"
	"sort"
	"time"
)

func StartClub(info *entity.ClubInfo, events []string) error {
	clients := make(map[string]entity.Client)
	tables := make([]entity.Table, info.NumberTables+1)

	tables[0] = entity.Table{
		IsUsing: true,
	}

	q := make([]string, 0)

	fmt.Printf("%s\n", info.OpenTime.Format("15:04"))

	prevTime := info.OpenTime

	for _, event := range events {

		if event == "" {
			continue
		}

		inputEvent, err := parser.ParseEvent(event, info)
		if err != nil {
			return fmt.Errorf("error parse event: %w", err)
		}

		if !utils.EventInWorkingTime(info.OpenTime, info.CloseTime, inputEvent.Time) && len(clients) > 0 {
			break
		} else if info.OpenTime.After(inputEvent.Time) {
			inputEvent.Time = inputEvent.Time.Add(24 * time.Hour)
		}
		fmt.Printf("%s\n", event)

		if !prevTime.Before(inputEvent.Time) && !prevTime.Equal(inputEvent.Time) && len(clients) > 0 {
			return fmt.Errorf("events must be sorted by time")
		}
		prevTime = inputEvent.Time

		switch inputEvent.ID {
		case entity.ClientCome:
			if _, exist := clients[inputEvent.ClientName]; exist {
				fmt.Printf("%s %d %s\n", inputEvent.Time.Format("15:04"),
					entity.EventFailed, "YouShallNotPass")
				break

			} else if !utils.EventInWorkingTime(info.OpenTime, info.CloseTime, inputEvent.Time) {
				fmt.Printf("%s %d %s\n", inputEvent.Time.Format("15:04"),
					entity.EventFailed, "NotOpenYet")
				break
			}
			clients[inputEvent.ClientName] = entity.Client{
				Status: entity.ClientCome,
			}

		case entity.ClientSeatTable:
			if _, exist := clients[inputEvent.ClientName]; !exist {
				fmt.Printf("%s %d %s\n", inputEvent.Time.Format("15:04"), entity.EventFailed, "ClientUnknown")
				break
			}
			if tables[inputEvent.TableNumber].IsUsing {
				fmt.Printf("%s %d %s\n", inputEvent.Time.Format("15:04"), entity.EventFailed, "PlaceIsBusy")
				break
			}

			if clients[inputEvent.ClientName].Status == entity.ClientSeatTable {
				client := clients[inputEvent.ClientName]
				periodSeat := inputEvent.Time.Sub(client.SeatTime)

				tables[client.Table].IsUsing = false
				tables[client.Table].UsedTime = tables[client.Table].UsedTime.Add(periodSeat)
				tables[client.Table].Profit += utils.CalculateProfit(info.PriceHour, periodSeat)
			}

			clients[inputEvent.ClientName] = entity.Client{
				Status:   entity.ClientSeatTable,
				Table:    inputEvent.TableNumber,
				SeatTime: inputEvent.Time,
			}
			tables[inputEvent.TableNumber].IsUsing = true

		case entity.ClientWaiting:
			if info.NumberTables <= len(q) {
				fmt.Printf("%s %d %s\n", inputEvent.Time.Format("15:04"), entity.ClientLeftClub, inputEvent.ClientName)
				delete(clients, inputEvent.ClientName)
				break
			}
			if slices.ContainsFunc(tables, func(t entity.Table) bool { return !t.IsUsing }) {
				fmt.Printf("%s %d %s\n", inputEvent.Time.Format("15:04"), entity.EventFailed, "ICanWaitNoLonger!")
				break
			}
			if clients[inputEvent.ClientName].Status == entity.ClientSeatTable {
				fmt.Printf("%s %d %s\n", inputEvent.Time.Format("15:04"), entity.EventFailed, "ClientAlreadyPlaying")
				break
			}

			client := clients[inputEvent.ClientName]
			client.Status = entity.ClientWaiting
			clients[inputEvent.ClientName] = client

			q = append(q, inputEvent.ClientName)
		case entity.ClientLeft:
			if _, exist := clients[inputEvent.ClientName]; !exist {
				fmt.Printf("%s %d %s\n", inputEvent.Time.Format("15:04"), entity.EventFailed, "ClientUnknown")
				break
			}
			client := clients[inputEvent.ClientName]

			if client.Status == entity.ClientSeatTable {
				periodSeat := inputEvent.Time.Sub(client.SeatTime)

				tables[client.Table].UsedTime = tables[client.Table].UsedTime.Add(periodSeat)
				tables[client.Table].Profit += utils.CalculateProfit(info.PriceHour, periodSeat)
			}

			delete(clients, inputEvent.ClientName)

			if len(q) == 0 {
				break
			}
			newClientName := q[0]
			q = q[1:]

			if clients[newClientName].Status != entity.ClientWaiting {
				fmt.Printf("%s %d %s\n", inputEvent.Time.Format("15:04"), entity.EventFailed, "ClientStatusNotWaiting")
				break
			}

			fmt.Printf("%s %d %s %d\n", inputEvent.Time.Format("15:04"), entity.ClientSeatAfterWaiting, newClientName, client.Table)

			clients[newClientName] = entity.Client{
				Status:   entity.ClientSeatTable,
				SeatTime: inputEvent.Time,
				Table:    client.Table,
			}
		default:
			log.Fatal("unknown event ID", inputEvent.ID)
		}
	}

	remainingClients := make([]string, 0, len(clients))

	for name := range clients {
		remainingClients = append(remainingClients, name)
	}

	sort.Strings(remainingClients)

	for _, name := range remainingClients {

		fmt.Printf("%s %d %s\n", info.CloseTime.Format("15:04"), entity.ClientLeftClub, name)

		client := clients[name]
		periodSeat := info.CloseTime.Sub(client.SeatTime)
		tables[client.Table].UsedTime = tables[client.Table].UsedTime.Add(periodSeat)
		tables[client.Table].Profit += utils.CalculateProfit(info.PriceHour, periodSeat)
		tables[client.Table].IsUsing = false

		delete(clients, name)
	}

	fmt.Printf("%v\n", info.CloseTime.Format("15:04"))
	for i, t := range tables[1:] {
		fmt.Printf("%d %d %v\n", i+1, t.Profit, t.UsedTime.Format("15:04"))
	}
	return nil
}
