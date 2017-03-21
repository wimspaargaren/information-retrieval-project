mapboxgl.accessToken = 'pk.eyJ1Ijoid2ltc3BhYXJnYXJlbiIsImEiOiJjajBjdGljODAwMDJrMnFud2x0c3M4YjZmIn0.D22MMfOi82i6LRjyntBwRw';
var map;
var geojson;
$(document).ready(function () {
    map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/streets-v9',
        zoom: 11,
        center: [4.899113, 52.372740]
    });

    var request = $.ajax({
        url: "http://localhost:8080/getpoints",
        method: "GET",
        dataType: "json"
    });
    request.done(function (msg) {
        geojson = msg;

    });
    request.fail(function (jqXHR, textStatus) {
        alert(textStatus);
    });

    map.on('load', function () {
        // Add points to the map
        map.addSource('point', {
            "type": "geojson",
            "data": geojson
        });

        map.addLayer({
            "id": "point",
            "type": "circle",
            "source": "point",
            'paint': {
                // make circles larger as the user zooms from z12 to z22
                'circle-radius': {
                    'base': 5,
                    'stops': [[12, 6], [22, 180]]
                },
                // color circles by ethnicity, using data-driven styles
                'circle-color': {
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
                }
            }
        });

    });

});





