
# Bus service

## Design patterns

- Visitor
    - bus/bus.go:VisitPassengers(..)
    - bus/bus.go:UpdatePassengers(..)

- Facade
    - internal/passenger.go:Passengers
    - internal/passenger.go:Add(..)
    - internal/passenger.go:Remove(..)
    - internal/passenger.go:Manifest(..)
    - internal/passenger.go:VisitPassengers(..)

- Observer
    - internal/busstop.go:NotifyBoardingIntent(..)

- Strategy
    - internal/buscompany.go:PriceCalculator

## Run

```bash
go run main.go
```

## Links

- https://medium.com/better-programming/learn-go-by-building-a-bus-service-6c11e7b81b92
