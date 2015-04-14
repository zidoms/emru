var gulp = require('gulp'),
	compass = require('gulp-compass');

var paths = {
	styles: 'styles/sass/*.sass'
};

function handleError(err) {
	console.log(err.toString());
	this.emit('end');
}

gulp.task('compass', function() {
	gulp.src(paths.styles)
		.pipe(compass({
			config_file: './config.rb',
			sass: 'styles/sass',
			css: 'styles'
		}))
		.on('error', handleError);
});

gulp.task('watch', function() {
	gulp.watch('styles/sass', ['compass']);
});

gulp.task('watch', ['compass', 'watch']);
gulp.task('default', ['compass']);
