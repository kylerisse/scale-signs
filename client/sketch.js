var clientState = "INIT";
var imgLogo;
var imgWifi;
var imgMap;

var testInt = 1;
var testCounter = 0; 


function setup() {
    frameRate(30);
    noCursor();

    createCanvas(windowWidth, windowHeight);

    imgLogo = loadImage("/img/logo-17x.png");
    imgWifi = loadImage("/img/wifi.png");
    imgMap = loadImage("/img/convention-center-map.png");
    loadSponsorImages();

    clientState = "READY";
}

function draw() {
    update();
    if (clientState === "READY") {
        background(255);
        imageMode(CENTER);
        image(imgLogo, 190, 100);
        image(imgWifi, windowWidth - 200, 100);
    }
    if (sponsorImages.length > 1) {
        image(sponsorImages[testInt], windowWidth - 130, windowHeight - 115);
        image(sponsorImages[testInt - 1], 130, windowHeight - 115);
    }
}

function windowResized() {
    resizeCanvas(windowWidth, windowHeight);
    centerCanvas();
}

function update() {
    testCounter++;
    if (testCounter > 300) {

        addSponsorImage();

        testInt++;
        testInt++;
        if (testInt >= sponsorImages.length) {
            testInt = 1;
        }
        testCounter = 0;
    }
}
