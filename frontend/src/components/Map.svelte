<style>
  #map {
    height: 750px;
    width: 100%;
    z-index: 0;
  }
</style>

<div
  id="map"
  data-markers="<%= @markers.to_json %>"
  data-mapbox-api-key="<%= ENV['MAPBOX_API_KEY'] %>" />
