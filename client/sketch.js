var clientState = "INIT";

var imgLogo;
var imgWifi;
var imgMap;

var corners = new sponsorcorners;

var timerFifteenSec = 0;
var timerFiveMin = 0;

function setup() {
    frameRate(30);
    noCursor();

    createCanvas(windowWidth, windowHeight);

    imgLogo = loadImage("/img/logo-17x.png");
    imgWifi = loadImage("/img/wifi.png");
    imgMap = loadImage("/img/convention-center-map.png");

    loadSponsorImages();

    testcounter = 290;
    clientState = "READY";
}

function draw() {
    update();
    if (clientState === "READY") {
        background(255);
    //    drawSchedule();
        corners.render();
        drawHeader();
    }
}

function windowResized() {
    resizeCanvas(windowWidth, windowHeight);
    centerCanvas();
}

function update() {
    timerFifteenSec++
    if (timerFifteenSec >= 450) {
        timerFifteenSec = 0;
        corners.update();
    }
    timerFiveMin++
    if (timerFiveMin >= 9000) {
        timerFiveMin = 0;
    }
}

function drawHeader() {
    imageMode(CENTER);
    image(imgLogo, 190, 100);
    image(imgWifi, windowWidth - 200, 100);
}
