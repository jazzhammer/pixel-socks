# pixel socks
2: runtime components: 

1. go: websocket server
2. nextjs: pixel-pushing client

# usage 
user picks a color, and may repick the color any number of times, to use as the pen for dropping pixels in the image, 
thus the instructions on the page: 

 - select a color
 - drop a pixel

when a pixel is dropped: 

 - the pixel data are saved in the database
 - the pixel data are broadcast to all sockets, including the pixel-dropper's socket

nb: this is a work in progress, watch the commits and try things out to see where the project is at. 
happy pixelating. 

