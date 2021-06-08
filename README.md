# jumper
Jump And Run Adventure game


# Notes:

## 2021-05-30

I want to create a jump-and-run game the working title has been set to "jumper".
For now I will focus on mechanics, engine, visualization becuase those are my storng points.
Visuals, animation, look-and-feel, gameplay, levels, story, come later.

Foundations:
- I want to write the game in Go, mostly because I want to learn the language.
- The above has been validated by finding [Ebiten](https://ebiten.org) which
    a library for writing cross-platform games in Go. 
    - Ebiten can run on web, ios, android and PC. The perfect combination for
      development and deployment.
    - Ebiten uses a stateless graphics model where the state is redrawn on 
      every render.
      (This will become important later.)
- I want to develop open source. I believe that this gives best progress for 
    now. 
- The game needs a strong community focus. Interactive multiplayer, leader-boards,
    custom levels are all a must.
- I want to build on the [Tiled editor](https://mapeditor.org). It is a great
    piece of tech.
  
I have read that multiplayer can always go through a server or p2p, but that
making it cheat-proof in p2p is difficult. P2p has the advantage of lower
latency, especially if the one server is located far away.

I have come up with the following architecture for a game engine to address and
make use of all of the above. Everything below is an initial design plan.

The game will consist of two parts:

- the world engine
- the game engine

### The world engine

The world progresses through ticks (perhaps 60/s). At any given time the entire
state of the world shall be representable in Tiled. There will be some global 
state
- PRNG state
- time in game

Some state per player
- score
- kill/die count
- inventory (health, ammo, ...)

Some state per object:
- velocity

All of this can be represented in Tiled. The advantages here are that the 
world engine can be developed with Tiled independently of the game engine.

The world engine takes the state of the world as one input, loaded from a file
or kept from the previous tick. The world engine then advances the world by
ticks. Nothing happens between ticks. The other input to the world engine is 
the vector of user actions `{time, action}` such as key-down. The world engine
will freely proagate the world until the action time, apply the action, then
be ready to advace on. The api to the world engine will be to give it actions
or to request the world state at a given time. The world engine will not use
system time in its calculations. This allows it to run faster than real time.

Any game state at time t1 will therefore be a function of two things, the world
state at an earlier time t0 and a vector of actions. Since the PRNG state is 
part of the world, the world engine is deterministic. This has important
conequences:
- for a given level, the action vector is sufficient to describe an entire 
played game.
- in multiplayer only actions need to be sent p2p, both game engines will then
 arrive at the same state.
- in multiplayer, both players can send their actions to the server where 
  another copy of the game engine can reproduce the entire game. The server
can verify scores and wins and identify cheaters.
- In single play, the action vector can be sent to the server 
  - to verify high-scores.
  - to allow replay, watching high-score runs...
  - adding a ghost for speed-running levels would be possible
- since action vectors don't fill much data space, the server could save a 
  very large number of replay games easily.
- games can easily be saved and resumed exactly, with everything, including
    random actions reoccuring precisely. (Note: change PRNG state on user
  actions to prevent providence.)
- since the world engine does not use wall-time it can potentially run very 
    fast, verifying an entire play session in a second. Any intermediate game
  state can be extracted from such a run, should it be desired.

### The game engine
The game engine has the task of rendering the game state. Here the fact that 
Ebiten re-renders on every frame is perfect. On every graphical tick, the 
game engine can ask the world engine for the state corresponding to the
wall-clock time and render it. For testing and development I can create any 
game state in Tiled and render it. It is the task of the game engine to 
allow for level selection, handle input scheme (touch, keyboard, ...) and
then to collect actions. These actions are then sent to the game engine.
(in multiplayer also to partner world engines). 

### Multi-player
In multi-player sessions it can happen that input actions are delayed and 
arrive in different orders and times at the game engines of different players.
Here the deteminism of the world engine is even more important. A buffer of
world states can be kept. If an action arrives from another player, that lies 
in the past, the world engine can take the last state from before that, apply
player B's action, and roll forward again to the current time. This may lead
to glitches, but truth and ordering are always preserved. To prevent too 
much glitching, the game engines should monitor latency to partner, and 
pause the game, showing "...synching..." if one player is getting too far 
behind (maybe 200ms). 

### Forum
- Hi-scores (time, kills, dies, points,) and full re-play of every hi-score run
- Browse level library
- social
    - channels
    - notify on new message/level/play of my levels
    - friends
    - voting
    - blocking
    - maker ranking
    - player ranking

### Development plan
The world engine and the game engine can be developed quite independently.
Both handle level state that is representable in Tiled.
- the world engine transforms it
  - I can create a state in Tiled, evolve it, and load it again in Tiled to
    see if it worked.
- the game engine displays it
    - I can create any state I want to see in Tiled.
I can switch between their development when I get bored with one or the other.

### More Notes on the world engine
- I want the geometry to be 90% encoded in the level. E.g. hit-boxes of moving
as well as stationary objects. This also means that e.g. the size of the player
  is not fixed in the engine (1 tile like RwK, 2 tiles like minecraft)
- Collectibles: Active and passive abilities, keys, switches, bonuses should use a plugin
 system to interact with "the laws of nature". The plugin system is not designed yet.
- Special blocks modify abilities and laws of nature. Do they use the plugin system?
- The laws of nature will not be variable.
  - Items fall (positive y direction). (Gravity might get a parameter)
  - Players have health and can die.
  - Players have a score.
  - Players have an inventory (probably cannot drop/interact with it directly)
  - Worlds may wrap in x, y, or both.
- The plugin system:
    - I have not decided yet if the plugins are scriptable outside the 
    code of the world engine itself.
    - I will certainly use strong OO design principles to make the 
      blocks/collectibles extendible with new additions
- The world engine will need to be able to emit events.
    - player won
    - player died
    - player gained ability that enables new input type
    - splash screen messages (e.g. "found a gun" or radio messages. The lingering will not be level state.)
- On the server the world engine will not interact with the game engine, but
with a server engine.
