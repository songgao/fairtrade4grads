var ctl = function() { };

ctl.prototype.getNames = function() {
    var self = this;
    $.getJSON('/api/list', function(data) {
        for(var i in data) {
            $('#name_list').append('<li class="list-group-item">' + data[i] + '</li>');
        }
    });
};

ctl.prototype.update_count = function() {
    $.get('/api/count', function(data) {
        $('#span_num').html(data);
    });
};

ctl.prototype.init = function() {
    $('#nav_signers').addClass('active');
    this.update_count();
    this.getNames();
};

$(window).load(ctl.prototype.init.bind(new ctl()));
