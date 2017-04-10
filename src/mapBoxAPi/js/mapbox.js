mapboxgl.accessToken = 'pk.eyJ1Ijoid2ltc3BhYXJnYXJlbiIsImEiOiJjajBjdGljODAwMDJrMnFud2x0c3M4YjZmIn0.D22MMfOi82i6LRjyntBwRw';
var map;
var geojson;
var geojsonPoly;
$(document).ready(function () {
    $('#showlegend').on('change', function (event) {
        if (event.target.checked) {
            $("#features")[0].style.display = "block";
        } else {
            $("#features")[0].style.display = "none";
        }
    })

    loadshp({
        url: 'http://localhost/voronoi.zip', // path or your upload file
        encoding: 'big5', // default utf-8
        EPSG: 3826 // default 4326
    }, function(geojson) {
        console.log(geojson)
         geojsonPoly = geojson;
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
                    'fill-color': {
                        property: 'sport-category',
                        type: 'categorical',
                        stops: [
                            ['soccer', 'red'],
                            ['fitness', 'green'],
                            ['running', 'blue'],
                            ['swimming', 'purple'],
                            ['fightingsport', 'yellow'],
                            ['cycling', 'orange'],
                            ['gymnastics', 'cyan'],
                            ['yoga', 'brown'],
                            ['hockey', 'white'],
                            ['bootcamp', 'pink']
                        ]
                    },
                    'fill-opacity': 0.8
                }
            });

    });


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
            console.log(geojsonPoly);
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
                    'fill-color': {
                        property: 'sport-category',
                        type: 'categorical',
                        stops: [
                            ['soccer', 'red'],
                            ['fitness', 'green'],
                            ['running', 'blue'],
                            ['swimming', 'purple'],
                            ['fightingsport', 'yellow'],
                            ['cycling', 'orange'],
                            ['gymnastics', 'cyan'],
                            ['yoga', 'brown'],
                            ['hockey', 'white'],
                            ['bootcamp', 'pink']
                        ]
                    },
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





