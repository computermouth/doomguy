## TODO

- pathing (busted)
- bun's walking shudder on random/zoom
- on sub_pos reset, set all other sub_pos to halfway??
- increase slime death animation timer
- non-battle event -- find weapon, find xp, find bed, story
- battle assist

- music
  - https://levelmusicmaxo.bandcamp.com/album/level-music-d (Lotus Pond loop for sleeping)
  - https://chipmusic.org/forums/topic/10093/game-dev-looking-for-music-read-this/
  - music from memory in regards to SDL (RWFromMem, Mix_LoadMUS_RW)
    - https://stackoverflow.com/questions/52500743/c-sdl2-mixer-play-wav-from-pointer

- non-60fps
```
while (delta > (16.7 * 2)) {
  ww_set_draw(0)
  process_state()
  delta -= 16.7
}
ww_set_draw(1)
process_state()

```
