
// standard library shit
#include <stdio.h>

// engine shit
#include "ww.h"

#include "enemy.h"

#define WINDOW_WIDTH  512
#define WINDOW_HEIGHT 384

void anim_end_callback(void * v){
	ww_sprite_t *s = (ww_sprite_t*)(v);
	s->active_animation++;

	if (s->active_animation > ENEMY_DEAD_INDEX )
		s->active_animation = 0;
}

int main( int argc, char * argv[] ) {
	
	// initialization
	if(ww_window_create(argc, argv, "Pixarray", WINDOW_WIDTH, WINDOW_HEIGHT)) {
		printf("Closing..\n");
		return 1;
	}
	
	ww_sprite_t * enemy = ww_new_sprite(ENEMY);
	enemy->animations[ENEMY_IDLE_INDEX  ].anim_end = anim_end_callback;
	enemy->animations[ENEMY_ATTACK_INDEX].anim_end = anim_end_callback;
	enemy->animations[ENEMY_HIT_INDEX   ].anim_end = anim_end_callback;
	enemy->animations[ENEMY_DEAD_INDEX  ].anim_end = anim_end_callback;
	
	// primary loop
	while(!ww_window_received_quit_event()) {
		
		// update events
		ww_window_update_events();
		
		ww_draw_sprite(enemy);
		
		// draw screen
		ww_window_update_buffer();
	}
	free(enemy);
	
	// cleanup and exit
	ww_window_destroy();
	return 0;
}
