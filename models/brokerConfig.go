package models

type BrokerConfig struct {
	Config        `bson:",inline"`
	Broker        string //`bson:"broker,omitempty"`
	SessionDomain string `bson:"sessionDomain,omitempty"`
}
