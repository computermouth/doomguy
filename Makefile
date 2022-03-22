
INCLUDES ?= -Isrc/ -Iimg/
PKG_CONFIG ?= pkg-config
SDL_LDFLAGS ?= `$(PKG_CONFIG) --libs sdl2`
SDL_CFLAGS ?= `$(PKG_CONFIG) --cflags sdl2`

F_LDFLAGS += $(SDL_LDFLAGS)
F_LDFLAGS += -flto
#~ F_LDFLAGS += -Wl,-z,norelro -Wl,--hash-style=gnu -Wl,--build-id=none
F_CFLAGS += -O3 -Wall -pedantic -std=gnu11
F_CFLAGS += -flto
#~ F_CFLAGS += -fno-stack-protector -fomit-frame-pointer -ffunction-sections -fdata-sections -Wl,--gc-sections
#~ F_CFLAGS += -fno-unwind-tables -fno-asynchronous-unwind-tables -fmerge-all-constants
F_CFLAGS += $(SDL_CFLAGS)

GAM_SRC = $(wildcard src/*.c)
GAM_OBJ = $(GAM_SRC:.c=.o)

IMG_DIR = img
IMG_YML = $(wildcard $(IMG_DIR)/*.yaml)
IMG_SRC = $(IMG_YML:.yaml=.c)
IMG_OBJ = $(IMG_SRC:.c=.o)

CONDIR = tools/concoord
CONCOORD = $(CONDIR)/concoord

all: | $(IMG_SRC) $(IMG_OBJ) $(GAM_OBJ)
	$(CC) -s -o main $(GAM_OBJ) $(IMG_OBJ) $(LDFLAGS) $(F_LDFLAGS)
#~ 	$(CC) -o main $(GAM_OBJ) $(IMG_OBJ) $(LDFLAGS) $(F_LDFLAGS)

$(IMG_SRC): | $(CONCOORD)

$(GAM_OBJ): | $(IMG_SRC)

$(IMG_DIR)/%.c: $(IMG_DIR)/%.yaml
	$(CONCOORD) $<

%.o: %.c
	$(CC) $(F_CFLAGS) $(INCLUDES) -c $< -o $@

$(CONCOORD):
	+$(MAKE) -C $(CONDIR)

memtest:
	valgrind --track-origins=yes --leak-check=yes ./main

opk:
	mksquashfs main assets/* facegame.opk -all-root -noappend -no-exports -no-xattrs -no-progress
