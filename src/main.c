
// standard library shit
#include <stdio.h>

// engine shit
#include "ww.h"

#define WINDOW_WIDTH  512
#define WINDOW_HEIGHT 384

int main( int argc, char * argv[] ) {
	
	// initialization
	if(ww_window_create(argc, argv, "Pixarray", WINDOW_WIDTH, WINDOW_HEIGHT)) {
		printf("Closing..\n");
		return 1;
	}
	
	// primary loop
	while(!ww_window_received_quit_event()) {
		
		// update events
		ww_window_update_events();
		
		// draw screen
		ww_window_update_buffer();
	}
	
	// cleanup and exit
	ww_window_destroy();
	return 0;
}
