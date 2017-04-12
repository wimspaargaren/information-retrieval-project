mapboxgl.accessToken = 'pk.eyJ1Ijoid2ltc3BhYXJnYXJlbiIsImEiOiJjajBjdGljODAwMDJrMnFud2x0c3M4YjZmIn0.D22MMfOi82i6LRjyntBwRw';

var map;


$(document).ready(function () {
    // to hide the different options based on the selection neighbourhoods/points
    $('#dataRepresentationSelect').on('change', function(event) {
        // points
        if (event.target.value == 1) {
            $("#dayfilter-div")[0].style.display = "block";
            $("#daypartfilter-div")[0].style.display = "block";
            $("#slider-div")[0].style.display = "none";
        } else {
            $("#dayfilter-div")[0].style.display = "none";
            $("#daypartfilter-div")[0].style.display = "none";
            $("#slider-div")[0].style.display = "block";
        }
    })
    
    
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

    var slider = document.getElementById('slider');
    var sliderValue = document.getElementById('slider-value');

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

        slider.addEventListener('input', function(e) {
            map.setPaintProperty('polygon', 'fill-opacity', parseInt(e.target.value, 10) / 100);
            sliderValue.textContent = e.target.value + '%';
        });
    });
});
