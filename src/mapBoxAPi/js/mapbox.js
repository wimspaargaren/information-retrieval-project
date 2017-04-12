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

    map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/light-v9',
        zoom: 13,
        center: [4.899113, 52.372740]
    });

    function getPolygons() {
        loadshp({
            url: 'http://localhost:8080/voronoi', // path or your upload file
            encoding: 'big5', // default utf-8
            EPSG: 3826 // default 4326
        }, function (geojson) {

            var requestPolygons = $.ajax({
                url: "http://localhost:8080/getpolygons",
                method: "GET",
                dataType: "json"
            });
            requestPolygons.done(function (msg) {
                console.log(msg);
                geojsonPoly = geojson;
                for(var cluster of msg.clusters) {
                    for(var id of cluster.ids) {
                        for(var f of geojsonPoly.features) {
                            if(f.properties.field_1 == id) {
                                f.properties.category = cluster.category;
                                // TODO: SET CATEGORY VAN CLUSTER, DEZE STAAT NU NIET IN DB
                            }
                        }
                    }
                }

                // // TODO: DIT MOET UITEINDELIJK WEG
                // for (var f of geojsonPoly.features) {
                //     if (f.properties.field_4 != "None") {
                //         f.properties.category = "running";
                //     }
                //     if (f.properties.field_5 != "None") {
                //         f.properties.category = "gymnastics";
                //     }
                //     if (f.properties.field_6 != "None") {
                //         f.properties.category = "cycling";
                //     }
                //     if (f.properties.field_7 != "None") {
                //         f.properties.category = "bootcamp";
                //     }
                //     if (f.properties.field_8 != "None") {
                //         f.properties.category = "fightingsport";
                //     }
                //     if (f.properties.field_9 != "None") {
                //         f.properties.category = "yoga";
                //     }
                //     if (f.properties.field_10 != "None") {
                //         f.properties.category = "soccer";
                //     }
                //     if (f.properties.field_11 != "None") {
                //         f.properties.category = "fitness";
                //     }
                //     if (f.properties.field_12 != "None") {
                //         f.properties.category = "swimming";
                //     }
                // }

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
                                ['bootcamp', colorBootcamp],
                                ['dancing', colorDancing]
                            ]
                        },
                        'fill-opacity': 0.5
                    }
                });

            });
            requestPolygons.fail(function (jqXHR, textStatus) {
                alert(textStatus);
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
        var layers = ['Soccer', 'Fitness', 'Running', 'Swimming', 'Fighting Sport', 'Cycling', 'Gymnastics', 'Yoga', 'Hockey', 'Bootcamp','Dancing'];
        var colors = [colorSoccer, colorFitness, colorRunning, colorSwimming, colorFighting, colorCycling, colorGymnastics, colorYoga, colorHockey, colorBootcamp,colorDancing];
        for (i = 0; i < layers.length; i++) {
            var color = colors[i];
            var item = document.createElement('div');
            var key = document.createElement('span');
            key.className = 'legend-key';
            key.style.backgroundColor = color;
            key.style.opacity = 0.5;

            var value = document.createElement('span');
            value.innerHTML = layers[i];
            item.appendChild(key);
            item.appendChild(value);
            $("#pd")[0].appendChild(item);
        }
    });
});





