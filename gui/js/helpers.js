App = {
	Models: {},
	Collections: {},
	Views: {}
};

template = function(id) {
	return _.template($('#' + id + '-template').html());
};

// dev
var gulp = require('gulp');
gulp.task('reload', function () {
if (location) location.reload();
});

gulp.watch('**/*', ['reload']);
