var ctl = function() {
    return this;
};

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
};

ctl.prototype.bindHandlers = function() {
    $('#btn_sign').click(this.sign);    
};

$(window).load(function() {
    new ctl().bindHandlers();
});
