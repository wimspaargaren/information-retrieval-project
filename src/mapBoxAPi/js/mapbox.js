mapboxgl.accessToken = 'pk.eyJ1Ijoid2ltc3BhYXJnYXJlbiIsImEiOiJjajBjdGljODAwMDJrMnFud2x0c3M4YjZmIn0.D22MMfOi82i6LRjyntBwRw';

var map;
var geojson;
var geojsonPoly;

var colorBootcamp = "#EF6C78";
var colorSoccer = "#FFFFEE"
var colorFitness = "#DADDF2"
var colorRunning = "#AABBDD"
var colorSwimming = "#8888AA"
var colorFighting = "#3F487F"
var colorCycling = "#223399"
var colorGymnastics = "#56141A"
var colorYoga = "#95222D"
var colorHockey = "#D53140"

$(document).ready(function () {
    $('#showlegend').on('change', function (event) {
        if (event.target.checked) {
            $("#features")[0].style.display = "block";
        } else {
            $("#features")[0].style.display = "none";
        }
    })

    map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/light-v9',
        zoom: 11,
        center: [4.899113, 52.372740]
    });

    function getPolygons() {
        loadshp({
            url: 'http://localhost/inforet/voronoi.zip', // path or your upload file
            encoding: 'big5', // default utf-8
            EPSG: 3826 // default 4326
        }, function (geojson) {

            geojsonPoly = geojson;

            for (var f of geojsonPoly.features) {
                if (f.properties.field_4 != "None") {
                    f.properties.category = "running";
                }
                if (f.properties.field_5 != "None") {
                    f.properties.category = "gymnastics";
                }
                if (f.properties.field_6 != "None") {
                    f.properties.category = "cycling";
                }
                if (f.properties.field_7 != "None") {
                    f.properties.category = "bootcamp";
                }
                if (f.properties.field_8 != "None") {
                    f.properties.category = "fightingsport";
                }
                if (f.properties.field_9 != "None") {
                    f.properties.category = "yoga";
                }
                if (f.properties.field_10 != "None") {
                    f.properties.category = "soccer";
                }
                if (f.properties.field_11 != "None") {
                    f.properties.category = "fitness";
                }
                if (f.properties.field_12 != "None") {
                    f.properties.category = "swimming";
                }

            }
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
                        property: 'category',
                        type: 'categorical',
                        stops: [
                            ['soccer', colorSoccer],
                            ['fitness', colorFitness],
                            ['running', colorRunning],
                            ['swimming', colorSwimming],
                            ['fightingsport', colorFighting],
                            ['cycling', colorCycling],
                            ['gymnastics', colorGymnastics],
                            ['yoga', colorYoga],
                            ['hockey', colorHockey],
                            ['bootcamp', colorBootcamp]
                        ]
                    },
                    'fill-opacity': 0.8
                }
            });
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
        var colors = [colorSoccer, colorFitness, colorRunning, colorSwimming, colorFighting, colorCycling, colorGymnastics, colorYoga, colorHockey, colorBootcamp];
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





