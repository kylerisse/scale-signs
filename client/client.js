class Client {

    constructor() {
        this.clientState = "INIT"; 
        this.imgLogo = loadImage("/img/logo-17x.png");
        this.imgWifi = loadImage("/img/wifi.png");
        this.imgMap = loadImage("/img/convention-center-map.png");
        
        this.sponsors = new Sponsors();
        this.sponsors.loadSponsorImages();

        this.testInt = 1;
        this.testCounter = 0;
        this.clientState = "READY";
    }

    tick() {
        this.testCounter++;
        if (this.testCounter > 300) {
            this.testInt++;
            this.testInt++;
            if (this.testInt >= this.sponsors.images.length) {
                this.testInt = 1;
            }
            this.testCounter = 0;
        }
        //console.log(this.testCounter + " " + this.testInt);
    }

    render() {
        if (this.clientState === "READY") {
            background(255);
            imageMode(CENTER)
            image(this.imgLogo, 190, 100);
            image(this.imgWifi, screen.width - 200, 100)
            if (this.sponsors.images.length > 1) {
                image(this.sponsors.images[testInt], screen.width - 130, screen.height - 115);
                image(this.sponsors.images[testInt - 1], 130, screen.height - 115);
            }
        }
    }

}