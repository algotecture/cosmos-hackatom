<script type="module">
  import SearchInput from "../components/SearchInput.svelte";
  import Map from "../components/Map.svelte";
</script>

<style>
  #search-flex {
    height: 100px;
    position: absolute;
    z-index: 3;
    display: flex;
    align-items: center;
    justify-content: center;
    left: 35vw;
  }
  .building-class,
  .tax-class,
  .address-icon {
    width: 155px;
    height: 155px;
  }

  .address-icon {
    padding: 20px;
  }

  .main-info-grid-container {
    display: grid;
    grid-template-columns: 155px auto 155px 155px;
    padding: 10px;
  }

  #banner-2 {
    background-color: rgb(251, 251, 251);
  }

  .main-container {
    position: relative;
  }

  #side-nav {
    width: 65px;
    height: 100%;
    position: absolute;
    background-color: #2e2672;
    z-index: 2;
  }

  #info-cards-container {
    width: 1100px;
    margin: 0 auto;
    position: relative;
    padding-top: 90px;
  }

  #main-info-card {
    background: white;
    width: 1100px;
    height: 190px;
    position: absolute;
    top: -95px;
    border-top: 5px solid #9ea4b6;
    border-radius: 0 0 9px 9px;
    box-shadow: 0 0 20px rgba(15, 15, 15, 0.1);
  }

  #info-cards-flex {
    display: flex;
    justify-content: space-between;
    padding-bottom: 120px;
  }

  .info-card {
    height: 320px;
    width: 23%;
    margin-top: 30px;
    background-color: white;
    border-top: 5px solid #3695bd;
    border-radius: 0 0 9px 9px;
    box-shadow: 0 0 20px rgba(15, 15, 15, 0.1);
  }
</style>

<div class="main-container">

  <div id="search-flex">
    <SearchInput />
  </div>
  <div id="side-nav" />

  <div id="banner-2">
    <Map />
  </div>

  <div id="info-cards-container">
    <div id="main-info-card">
      <div class="main-info-grid-container">
        <img
          alt="address"
          src="address-icon.png"
          class="grid-item address-icon" />
        <div class="grid-item address">
          <h1>Address</h1>
          <p>Address</p>
        </div>
        <div class="grid-item building-class">building class</div>
        <div class="grid-item tax-class">tax class</div>
      </div>
    </div>
    <div id="info-cards-flex">
      <div class="info-card" />
      <div class="info-card" />
      <div class="info-card" />
      <div class="info-card" />
    </div>
  </div>
</div>
