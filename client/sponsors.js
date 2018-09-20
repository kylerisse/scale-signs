var sponsorImageList = [];
var sponsorImages = [];

function loadSponsorImages() {
    loadJSON("/api/sponsors/images", function(list) {
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
