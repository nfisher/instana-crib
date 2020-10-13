"use strict";

function heatmap(url, div, count) {
    return function() {
        let x=d3.select(div).node().getBoundingClientRect().width;

        // set the dimensions and margins of the graph
        let margin = {top: 5, right: 30, bottom: 55, left: 35},
            width = x - margin.left - margin.right,
            height = 150 - margin.top - margin.bottom;

        //Read the data
        d3.csv(url, function(data) {
            d3.select(div)
                .select("svg")
                .remove();
            // append the svg object to the body of the page
            let svg = d3.select(div)
                .append("svg")
                .attr("width", width + margin.left + margin.right)
                .attr("height", height + margin.top + margin.bottom)
                .append("g")
                .attr("transform",
                    "translate(" + margin.left + "," + margin.top + ")");
            let groups = {};
            let subGroups = {};
            let upper = 0;
            for (let i = 0; i < data.length; i++) {
                let g = data[i]["group"];
                let s = data[i]["variable"];
                groups[g] = true;
                subGroups[s] = true;
            }
            for (let i = 0; i < Object.keys(subGroups).length; i++) {
                upper += parseInt(data[i]["value"]);
            }
            // Labels of row and columns
            let myGroups = Object.keys(groups);
            let myVars = Object.keys(subGroups);

            // Build X scales and axis:
            let x = d3.scaleBand()
                .range([ 0, width])
                .domain(myGroups)
                .padding(0.01);
            svg.append("g")
                .attr("transform", "translate(0," + height + ")")
                .call(d3.axisBottom(x).tickValues(x.domain().filter(function(d,i){ return !(i%2)})))
                .selectAll("text")
                .style("text-anchor", "end")
                .attr("dx", "-.8em")
                .attr("dy", ".15em")
                .attr("transform", "rotate(-65)");

            // Build X scales and axis:
            let y = d3.scaleBand()
                .range([ height, 0 ])
                .domain(myVars)
                .padding(0.01);
            svg.append("g")
                .call(d3.axisLeft(y).tickValues(y.domain().filter(function(d,i){ return !(i%5)})));

            d3.select(count)
                .text("" + upper);
            // Build color scale
            let myColor = d3.scaleLinear()
                .range(["white", "#eee", "#990000"])
                .domain([0, 1, upper]);
            if (upper === 1) {
               myColor = d3.scaleLinear()
                   .range(["white", "#990000"])
                   .domain([0, 1]);
            }
            svg.selectAll()
                .data(data, function(d) {return d.group+':'+d.variable;})
                .enter()
                .append("rect")
                .attr("x", function(d) { return x(d.group) })
                .attr("y", function(d) { return y(d.variable) })
                .attr("width", x.bandwidth() )
                .attr("height", y.bandwidth() )
                .style("fill", function(d) { return myColor(d.value)} )
        })
    }
}

function onResizeInterval(fn, interval) {
    fn();
    window.addEventListener('resize', fn);
    setInterval(fn, interval);
}

function spark(url, id) {
    return function() {
        d3.json(url, function(data) {
            const WIDTH = 180;
            const HEIGHT = 20;
            const DATA_COUNT = data.values.length;
            const BAR_WIDTH = (WIDTH - DATA_COUNT) / DATA_COUNT;
            const x = d3.scaleLinear().domain([0, DATA_COUNT]).range([0, WIDTH]);
            const y = d3.scaleLinear().domain([0, d3.max(data.values)]).range([HEIGHT, 0]);

            let pctl99 = Math.round(d3.quantile(data.values, 0.99));
            d3.select(id + "99")
                .text(pctl99);

            let last = Math.round(data.values[data.values.length - 1]);
            d3.select(id + "Last")
                .text(last);

            let max = Math.round(d3.max(data.values));
            d3.select(id + "Max")
                .text(max);

            d3.select(id)
                .select("svg")
                .remove();
            const svg = d3.select(id).append("svg")
                .attr("width", WIDTH)
                .attr("height", HEIGHT)
                .append("g");
            svg.selectAll(".bar").data(data.values)
                .enter()
                .append("rect")
                .attr("class", "bar")
                .attr("x", (d, i) => x(i))
                .attr("y", d => HEIGHT - y(d))
                .attr("width", BAR_WIDTH)
                .attr("height", d => y(d))
                .attr("fill", "#ccc");
        });
    };
}

function main() {
    let adProcessor = heatmap("heatmap_data?metric=metrics.gauges.KPI.incoming.span_messages.error_rate&entity=appdataProcessor", "#g_ad_processor_dropping", "#g_ad_processor_dropping_count");
    let adWriter = heatmap("heatmap_data?metric=metrics.gauges.KPI.incoming.raw_spans.error_rate&entity=appdataWriter", "#g_ad_writer_dropping", "#g_ad_writer_dropping_count");
    let cpuSys = heatmap("heatmap_data?metric=cpu.sys&entity=host", "#g_cpu_sys", "#g_cpu_sys_count");
    let cpuUser = heatmap("heatmap_data?metric=cpu.user&entity=host", "#g_cpu_user", "#g_cpu_user_count");
    let cpuWait = heatmap("heatmap_data?metric=cpu.wait&entity=host", "#g_cpu_wait", "#g_cpu_wait_count");
    let fillerDropping = heatmap("heatmap_data?metric=metrics.gauges.KPI.incoming.raw_messages.error_rate&entity=filler", "#g_filler_dropping", "#g_filler_dropping_count");
    let fillerSpark = spark("ts_sum?entity=filler&metric=metrics.gauges.com.instana.filler.service.snapshot.OnlineSnapshotsLimit.online-snapshots-count", "#sparkFiller");
    let adProcessorSpark = spark("ts_sum?entity=appdataProcessor&metric=metrics.meters.KPI.incoming.span_messages.calls", "#sparkProcessor");
    let adWriterSpark = spark("ts_sum?entity=appdataWriter&metric=metrics.meters.KPI.incoming.raw_spans.calls", "#sparkWriter");

    onResizeInterval(fillerSpark, 250);
    onResizeInterval(adProcessorSpark, 250);
    onResizeInterval(adWriterSpark, 250);
    onResizeInterval(adProcessor, 250);
    onResizeInterval(adWriter, 250);
    onResizeInterval(cpuSys, 250);
    onResizeInterval(cpuUser, 250);
    onResizeInterval(cpuWait, 250);
    onResizeInterval(fillerDropping, 250);
}

main();
