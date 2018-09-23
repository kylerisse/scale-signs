var clientState = "INIT";

var imgLogo;
var imgWifi;
var imgMap;

function setup() {
    frameRate(30);
    noCursor();

    createCanvas(windowWidth, windowHeight);

    imgLogo = loadImage("/img/logo-17x.png");
    imgWifi = loadImage("/img/wifi.png");
    imgMap = loadImage("/img/convention-center-map.png");

    initSponsors();

    clientState = "READY";
}

function draw() {
    update();
    if (clientState === "READY") {
        background(223, 202, 150);
        sponsorsRender();
        drawHeader();
    }
}

function windowResized() {
    resizeCanvas(windowWidth, windowHeight);
    centerCanvas();
}

function update() {
    sponsorsTick();
}

function drawHeader() {
    fill(255);
    stroke(255);
    rect(0, 0, windowWidth, 190)
    imageMode(CENTER);
    image(imgWifi, windowWidth - 200, 100);
    image(imgLogo, 190, 100);
}
