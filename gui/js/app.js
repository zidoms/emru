(function() {

window.App = {
	Models: {},
	Collections: {},
	Views: {},
	gui: require('nw.gui')
};

window.template = function(id) {
	return _.template($('#' + id + 'Template').html());
};

})();
