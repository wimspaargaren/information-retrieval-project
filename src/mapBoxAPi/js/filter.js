
function filterOnDay(number) {
    var day = getDayFromNumber(number);
    map.removeSource("point")

    if(map.getLayer(day) !== undefined) {
        var visibility = map.getLayoutProperty(day, 'visibility');

        if (visibility === 'visible') {
            map.setLayoutProperty(day, 'visibility', 'none');
            this.className = '';
        } else {
            this.className = 'active';
            map.setLayoutProperty(day, 'visibility', 'visible');
        }
    } else {
        var resJson = {features: [], type: "FeatureCollection"};
        for(var i=0; i<geojson.features.length; i++) {
            if(geojson.features[i].properties.day == day) {
                resJson.features.push(geojson.features[i]);
            }
        }

        map.addSource('point', {
            "type": "geojson",
            "data": resJson
        });

        map.addLayer({
            "id": day,
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
                        ['Voetbal', 'red'],
                        ['Tennis', 'green'],
                        ['hardlopen', 'blue']
                    ]
                }
            }
        });
    }
}

function getDayFromNumber(number) {
    switch(number) {
        case 1 : return "Monday";
        case 2 : return "Tuesday";
        case 3 : return "Wednesday";
        case 4 : return "Thursday";
        case 5 : return "Friday";
        case 6 : return "Saturday";
        case 7 : return "Sunday";
        default: return "Monday";
    }
}