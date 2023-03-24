package xatu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventFilter(t *testing.T) {
	testConfig := &EventFilterConfig{
		EventNames: []string{},
	}

	filter, err := NewEventFilter(testConfig)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, filter)
}

func TestNewEventFilter_InvalidName(t *testing.T) {
	testConfig := &EventFilterConfig{
		EventNames: []string{"InvalidEventName"},
	}

	_, err := NewEventFilter(testConfig)

	assert.Error(t, err)
}

func TestEventFilter_Apply(t *testing.T) {
	testEvents := []*DecoratedEvent{{
		Event: &Event{
			Name: Event_BEACON_API_ETH_V1_DEBUG_FORK_CHOICE,
		},
	}, {
		Event: &Event{
			Name: Event_BEACON_API_ETH_V1_EVENTS_ATTESTATION,
		},
	}, {
		Event: &Event{
			Name: Event_BEACON_API_ETH_V1_EVENTS_BLOCK,
		},
	}}

	testConfig := &EventFilterConfig{
		EventNames: []string{Event_BEACON_API_ETH_V1_DEBUG_FORK_CHOICE.String()},
	}
	filter, _ := NewEventFilter(testConfig)

	filteredEvents, err := filter.Apply(testEvents)
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, filteredEvents, 1)
	assert.Equal(t, testEvents[0], filteredEvents[0])
}

func TestEventFilter_ShouldBeDropped(t *testing.T) {
	testEvent := &DecoratedEvent{
		Event: &Event{
			Name: Event_BEACON_API_ETH_V1_DEBUG_FORK_CHOICE,
		},
	}

	testConfig := &EventFilterConfig{
		EventNames: []string{Event_BEACON_API_ETH_V1_EVENTS_FINALIZED_CHECKPOINT.String()},
	}
	filter, _ := NewEventFilter(testConfig)

	shouldBeDropped, err := filter.ShouldBeDropped(testEvent)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, shouldBeDropped)
}

func TestEventFilter_AllowEverythingWhenEmpty(t *testing.T) {
	events := []*DecoratedEvent{}

	for _, eventName := range Event_Name_value {
		events = append(events, &DecoratedEvent{
			Event: &Event{
				Name: Event_Name(eventName),
			},
		})
	}

	emptyConfig := &EventFilterConfig{}
	filter, _ := NewEventFilter(emptyConfig)

	filteredEvents, err := filter.Apply(events)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, events, filteredEvents)
}
