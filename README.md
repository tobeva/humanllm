# Human LLM

An LLM made of people.

An experimental web app where humans pick the next word through a collaborative
creative process. Each player can vote for the next word, so every player can
influence the output. The hope is it resembles using a Ouija board in two ways:

1. The output emerges the whole group, no one person is in changer.
2. Player are sometimes surprised and impressed by the output.

# Main Page

Minimal mostly-text page with a title and a list of open sessions. Each session
has a name you can click on, plus some attributes underneath such as:

  * The number of current players
  * How long the session has been going, or how long until the timer ends
  * How many words have been chosen so far
  * Any rule modifiers for that session
  * Other ideas?

# Session Page

The session page has two main elements:

1. An in-progress sentence or paragraph in a large font, ending with an
    underscore indicating the next word to be chosen.
2. A "word cloud" or some sort of list of possible next words where you can vote
    for a specific word.

A count-down indicator lets the player know when their vote will be locked in
and the next word chosen. Possible variations on how to vote for a word:

  * You can select only a single word, if you click again it changes your vote.
  * You can select your top N favorite words
  * You can keep clicking the same word to increase your vote.
  * You can keep clicking on as many words as you want.
  * Other ideas?

# Architecture

See [ARCHITECTURE.md](ARCHITECTURE.md)
