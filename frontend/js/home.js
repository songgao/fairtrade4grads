var ctl = function() { };

ctl.prototype.sign = function() {
    var self = this;
    signature = {
        'name': $('#input_name').val(),
        'au_username': $('#input_username').val()
    };
    $('#btn_sign').text('Sending Request...');
    $('button').prop('disabled', true);
    $.post( '/api/sign', JSON.stringify(signature), ctl.prototype.sign_successful.bind(self));
};

ctl.prototype.sign_successful = function() {
    $('#btn_sign').text('Succesfully Signed');
    $('button').prop('disabled', true);
    this.update_count();
};

ctl.prototype.update_count = function() {
    $.get('/api/count', function(data) {
        $('#span_num').html(data);
    });
};

ctl.prototype.init = function() {
    $('#btn_sign').click(ctl.prototype.sign.bind(this));    
    $('#nav_home').addClass('active');
    this.update_count();
};

$(window).load(ctl.prototype.init.bind(new ctl()));
