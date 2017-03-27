mapboxgl.accessToken = 'pk.eyJ1Ijoid2ltc3BhYXJnYXJlbiIsImEiOiJjajBjdGljODAwMDJrMnFud2x0c3M4YjZmIn0.D22MMfOi82i6LRjyntBwRw';
var map;
var geojson;
var geojsonPoly;
$(document).ready(function () {
    map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/light-v9',
        zoom: 11,
        center: [4.899113, 52.372740]
    });

    function getPolygons() {
        var requestPolygons = $.ajax({
            url: "http://localhost:8080/getpolygons",
            method: "GET",
            dataType: "json"
        });
        requestPolygons.done(function (msg) {
            geojsonPoly = msg;
            map.addSource('polygon', {
                "type": "geojson",
                "data": geojsonPoly
            });
            map.addLayer({
                "id": "polygon",
                "type": "fill",
                "source": "polygon",
                'layout': {},
                'paint': {
                    'fill-color': '#088',
                    'fill-opacity': 0.8
                }
            });

        });
        requestPolygons.fail(function (jqXHR, textStatus) {
            alert(textStatus);
        });
    }

    function getPoints() {
        var requestPolygons = $.ajax({
            url: "http://localhost:8080/getpoints",
            method: "GET",
            dataType: "json"
        });
        requestPolygons.done(function (msg) {
            geojson = msg;
        });
        requestPolygons.fail(function (jqXHR, textStatus) {
            alert(textStatus);
        });
    }
    getPolygons();
    getPoints();

    map.on('load', function () {

        map.getCanvas().style.cursor = 'default';
        var layers = ['Soccer', 'Fitness', 'Running', 'Swimming', 'Fighting Sport', 'Cycling', 'Gymnastics', 'Yoga', 'Hockey', 'Bootcamp'];
        var colors = ['red', 'green', 'blue', 'purple', 'yellow', 'orange', 'cyan', 'brown', 'white', 'pink'];
        for (i = 0; i < layers.length; i++) {
            var color = colors[i];
            var item = document.createElement('div');
            var key = document.createElement('span');
            key.className = 'legend-key';
            key.style.backgroundColor = color;

            var value = document.createElement('span');
            value.innerHTML = layers[i];
            item.appendChild(key);
            item.appendChild(value);
            $("#pd")[0].appendChild(item);
        }
    });
});





