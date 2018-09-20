var client; 

function preload() {
    client = new Client();
}

function setup() {
    frameRate(30);
    noCursor();
    createCanvas(screen.width, screen.height);
}

function draw() {
    client.tick();
    client.render();
}


