package main

import (
  "os"
  "encoding/json"
  "errors"
)

type CardEffect func (*Board)

type Card struct {
  id string
  name string
  effect CardEffect
}

cardEffectMap = map[string]CardEffect

func (this *Card) loadCard(id string) (*Card, error) {
  file, err := os.Open("cards/" + id)
  if err != nill {
    return nil, err
  }
  
  var jsonMap map[string]interface{}
  decoder := json.Decoder(file)
  err := json.Decode(file, &jsonMap)
  if err != nil {
    return nil, err
  }

  card := Card{id, "", nil}
  name, ok = json.map['name'].(string)
  if !ok {
    return nil, errors.New("Bad name")
  }
  card.name = name

  effect, ok := cardEffectMap[name]
  if !ok {
    return nil, errors.New("Bad effect name")
  }
  card.effect = effect

  return &card
}

type Deck struct {
  cards []*Card
}

func (this *Deck) UnmarshalJSON(b []byte) error {
  var parsed map[string]interface{}
  this.cards := make([]*Card)

  err := json.Unmarshal(b, &parsed)
  if err != nil {
    return err
  }
  
  cards, ok := parsed['cards'].([]string)
  if !ok {
    return errors.New("Invalid json")
  }

  for _, card_s := range parsed['cards'] {
    card, err := loadCard(card_s)
    this.cards = append(this.cards, card)
  }
}

func (this *Deck) shift() *Card {
  card := this.cards[0]
}

func (this *Deck) getCard(i int) *Card {
  return this.cards[i]
}

func (this *Deck) deleteCard(i int) *Card {
  card := this.cards[i]
  this.cards = append(a[:i], a[i+1:]...)
  return card
}

func (this *Deck) shuffle() {
  Shuffle(this.cards)
}

func loadDeck(id string) (*Deck, error) {
  file, err := os.Open("decks/" + id)
  if err != nill {
    return nil, err
  }

  decoder := json.NewDecoder(file)
  deck := Deck{}
  err := decoder.Decode(&deck)
  if err != nil {
    return nil, err
  }
  return deck, nil
}

func Shuffle(interface{} []int) {
    for i := range a {
        j := rand.Intn(i + 1)
        a[i], a[j] = a[j], a[i]
    }
}
