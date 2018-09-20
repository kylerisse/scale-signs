var clientState = "INIT";
var imgLogo;
var imgWifi;
var imgMap;

var testInt = 1;
var testCounter = 0; 


function setup() {

    frameRate(30);
    noCursor();
    createCanvas(screen.width, screen.height);

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
        imageMode(CENTER)
        image(imgLogo, 190, 100);
        image(imgWifi, screen.width - 200, 100)
    }
    if (sponsorImages.length > 1) {
        image(sponsorImages[testInt], screen.width - 130, screen.height - 115);
        image(sponsorImages[testInt - 1], 130, screen.height - 115);
    }
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
