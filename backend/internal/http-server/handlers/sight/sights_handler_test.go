package sight_test

// import (
// 	"context"
// 	"testing"

// 	"homework_ipl/internal/entities"
// 	"homework_ipl/internal/http-server/handlers/sight"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetSights(t *testing.T) {
// 	handler := &sight.SightsHandler{}

// 	ctx := context.Background()

// 	resp, err := handler.GetSights(ctx, entities.Sight{})
// 	if err != nil {
// 		t.Fatalf("Failed to get sights: %v", err)
// 	}

// 	assert.NotEmpty(t, resp.Sight)

// 	expectedFirstSight := entities.Sight{
// 		ID:          1,
// 		Rating:      4.3434,
// 		Name:        "Парижская башня",
// 		Description: "Самая высокая башня в мире.",
// 		City:        "Париж",
// 		Url:         "public/1.jpg",
// 	}
// 	assert.Equal(t, expectedFirstSight, resp.Sight[0])
// }
