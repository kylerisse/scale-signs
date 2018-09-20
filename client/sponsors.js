
class Sponsors {

    constructor() {
        this.images = [];
    }

    loadSponsorImages() {
        loadJSON("/api/sponsors/images", function(list) {
            console.log("test");
            this.images = list;
        });
        console.log(this.images);
        console.log(this.images[0]);
        for (let i = 0; i < this.images; i++) {
            console.log("/img/sponsors/" + this.images[0]);
            this.images.push(loadImage("/img/sponsors/" + this.images[i]));
        }
    } 

}