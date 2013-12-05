var ctl = function() { return this; };

ctl.prototype.init = function() {
    $('#nav_faq').addClass('active');
};

$(window).load(ctl.prototype.init.bind(new ctl()));
