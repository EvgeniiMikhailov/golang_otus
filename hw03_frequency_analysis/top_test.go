package hw03_frequency_analysis //nolint:golint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func contains(words []string, target string) bool {
	for _, word := range words {
		if word == target {
			return true
		}
	}
	return false
}

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		assert.Len(t, Top10(""), 0)
	})

	t.Run("simple text line of Text", func(t *testing.T) {
		text := "simple text line of Text"
		get := Top10(text)
		assert.Len(t, get, 4)
		assert.Equal(t, get[0], "text")
	})

	t.Run("simple text line of Text", func(t *testing.T) {
		text := "simple text line of Text"
		get := Top10(text)
		assert.Len(t, get, 4)
		assert.Equal(t, get[0], "text")
	})

	t.Run("only space in text", func(t *testing.T) {
		text := " "
		get := Top10(text)
		assert.Len(t, get, 0)
	})

	t.Run("only punctuation in text", func(t *testing.T) {
		text := "!"
		get := Top10(text)
		assert.Len(t, get, 0)
	})

	t.Run("only punctuation and spaces in text", func(t *testing.T) {
		text := " ! "
		get := Top10(text)
		assert.Len(t, get, 0)
	})

	t.Run("two consecutive spaces in text", func(t *testing.T) {
		text := "a  b"
		get := Top10(text)
		assert.Len(t, get, 2)
	})

	t.Run("Text line of, Text with punctuation symbols", func(t *testing.T) {
		text := "Text line of, Text with! punctuation? symbols"
		get := Top10(text)
		assert.Len(t, get, 6)
		assert.Equal(t, get[0], "text")
	})

	t.Run("test 11 words", func(t *testing.T) {
		text := "a a b b c c d d e e f f g g h h i i j j k"
		get := Top10(text)
		assert.Len(t, get, 10)
		assert.False(t, contains(get, "k"))
	})

	t.Run("test order", func(t *testing.T) {
		text := "a a a a a a a a a a b b b b b b b b b c c c c c c c c d d d d d d d e e e e e e f f f f f g g g g h h h i i j"
		get := Top10(text)
		assert.Len(t, get, 10)
		assert.Equal(t, get, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"})
	})

	t.Run("test unicode", func(t *testing.T) {
		text := "Привет мир привет!"
		get := Top10(text)
		assert.Len(t, get, 2)
		assert.Equal(t, get, []string{"привет", "мир"})
	})

	t.Run("positive test", func(t *testing.T) {
		//t.Skip("Skipping test while in development")
		if taskWithAsteriskIsCompleted {
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
			assert.Subset(t, expected, Top10(text))
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			assert.ElementsMatch(t, expected, Top10(text))
		}
	})
}
