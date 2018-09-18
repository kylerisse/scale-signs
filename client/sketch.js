var clientState;
var imgLogo;
var imgWifi;
var imgMap;
var imgAllSponsors;

var testInt;
var testCounter;

function setup() {
    clientState = "INIT";
    frameRate(30);
    noCursor();
    createCanvas(screen.width, screen.height);
    imgLogo = loadImage("/img/logo-17x.png");
    imgWifi = loadImage("/img/wifi.png");
    imgMap = loadImage("/img/convention-center-map.png");
    clientState = "READY";
    imgAllSponsors = loadSponsorImages();

    testInt = 1;
    testCounter = 0;

}

function draw() {
    if (clientState === "READY") {
        background(255);
        imageMode(CENTER)
        image(imgLogo, 190, 100);
        image(imgWifi, screen.width - 200, 100)
        image(imgAllSponsors[testInt], screen.width - 130, screen.height - 115);
        image(imgAllSponsors[testInt - 1], 130, screen.height - 115);
        testCounter++;
        if (testCounter > 300) {
            testInt++;
            testInt++;
            if (testInt >= imgAllSponsors.length) {
                testInt = 1;
            }
            testCounter = 0;
        } 
    }

}
