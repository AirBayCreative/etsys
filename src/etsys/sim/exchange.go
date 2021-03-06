package sim

import (
	. "etsys"
)

type SimulatedExchange struct {
	markets  map[string]*SimulatedMarket
	Tradelog chan *Trade
	Orderlog chan *OrderState
}

func (se *SimulatedExchange) Run() {
	for m := range se.markets {
		se.markets[m].Run()
	}
}

func (se *SimulatedExchange) AttachMarket(m *SimulatedMarket) {
	se.markets[m.Ticker] = m
}

func (se *SimulatedExchange) GetMarket(ticker string) *SimulatedMarket {
	return se.markets[ticker]
}
func (se *SimulatedExchange) GetTickers() []string {
	ts := make([]string, 0)
	for t := range se.markets {
		ts = append(ts, t)
	}
	return ts
}
func (se *SimulatedExchange) SendOrder(o *Order) {
	m := se.markets[o.Ticker]
	if m == nil {
		panic("wrong exchange")
	}
	m.OrderReciever <- o
}

func MakeSimulatedExchange(orderlog chan *OrderState, tradelog chan *Trade) *SimulatedExchange {
	se := &SimulatedExchange{
		markets:  make(map[string]*SimulatedMarket),
		Tradelog: tradelog,
		Orderlog: orderlog,
	}
	return se
}

func MakeSomeSimulatedExchange(orderlog chan *OrderState, tradelog chan *Trade) *SimulatedExchange {
	se := MakeSimulatedExchange(orderlog, tradelog)
	se.AttachMarket(MakeSimulatedMarket("A", se.Orderlog, se.Tradelog))
	se.AttachMarket(MakeSimulatedMarket("B", se.Orderlog, se.Tradelog))
	return se
}
