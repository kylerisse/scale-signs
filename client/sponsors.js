var sponsorImageList = [];
var sponsorImages = [];

function loadSponsorImages() {
    loadJSON("/api/sponsors", function(list) {
        sponsorImageList = list;
    });
} 

function addSponsorImage(string) {
    if (sponsorImageList.length != sponsorImages.length) {
        sponsorImages = [];
        for (let i = 0; i < sponsorImageList.length; i++) {
            sponsorImages.push(loadImage("/img/sponsors/" + sponsorImageList[i]));
        }
    }
}

class sponsorcorners {

    constructor() {
        var leftImg;
        var rightImg;
    }

    render() {
        if (sponsorImages.length > 1) {
            imageMode(CORNERS)
            draw(sponsorImages[this.leftImg], sponsorsImages[this.leftImg].width + 150, windowHeight - sponsorsImage[this.leftImg].height - 150)
            draw(sponsorImages[this.rightImg], windowWidth - sponsorsImage[this.rightImg].width - 150, windowHeight - sponsorImage[this.rightImg].height)
        }
    }

    update() {
        if (sponsorImages.length > 2 && timerFifteenSec == 0) {
            leftImg = rand(sponsorImages.length - 1)
            rightTmp = rand(sponsorImages.length - 1)
            while (leftImg == rightTmp) {
                rightTmp = rand(sponsorImages.length - 1)
            }
            rightImg = rightTmp
        }
    }

}