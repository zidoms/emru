var request = require('request');

App = {
	Models: {},
	Collections: {},
	Views: {}
};

template = function(id) {
	return _.template($('#' + id + '-template').html());
};

// overriding backbone sync to use request library
Backbone.sync = function(method, model, options) {
	var type = (!options.type) ? methodMap[method] : options.type;

	// Default options, unless specified.
	_.defaults(options || (options = {}), {});

	// Default JSON-request options.
	var params = {type: type, dataType: 'json', headers: {}};

	// Ensure that we have a URL.
	if (!options.url)
		params.url = _.result(model, 'url') || urlError();
	else
		params.url = options.url;

	// Ensure that we have the appropriate request data.
	if (!options.data && model && (method === 'create' || method === 'update' || method === 'patch')) {
		params.contentType = 'application/json';
		params.data = model.toJSON();
	}

	// Don't process data on a non-GET request.
	if (params.type !== 'GET')
		params.processData = false;

	request({ url: params.url, json: true, method: params.type, headers: params.headers, body: params.data }, function (err, result, body) {
		if (err)
			return options.error(err);

		return options.success(body);
	});
};

// Map from CRUD to HTTP for our default `Backbone.sync` implementation.
var methodMap = {
	'create': 'POST',
	'update': 'PUT',
	'patch': 'PATCH',
	'delete': 'DELETE',
	'read': 'GET'
};

function isRTL(str) {
	var regex = new RegExp('[\u0591-\u07FF\uFB1D-\uFDFD\uFE70-\uFEFC]', 'g');
	str = str.replace(/\s/g, '');

	res = str.match(regex);
	if (res !== null && res.length/str.length > 0.5)
		return true;
	return false;
}

// dev
var gulp = require('gulp');
gulp.task('reload', function () {
	if (location) location.reload();
});

gulp.watch('**/*', ['reload']);
