var app = new Framework7();
var $$ = Dom7;

$$('.toggle').on('change', function (e) {
    const that = $$(e.target);
    const gpio = that.val();
    const checked = that.prop('checked');
    app.request.json(`/toggle/${gpio}/${checked?1:0}`, function (data) {
        console.log(data);
    });
  });