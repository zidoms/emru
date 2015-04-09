var gulp = require('gulp'),
	compass = require('gulp-compass'),
	htmlmin = require('gulp-htmlmin'),
	uglify = require('gulp-uglify'),
	concat = require('gulp-concat');

var paths = {
	js: 'js/*.js',
	libjs: [
		'js/lib/underscore.js',
		'js/lib/jquery.js',
		'js/lib/backbone.js',
		'dist/app.js'
	],
	styles: 'styles/*.sass',
	htmls: '*.html',
	dist: 'dist'
};

function handleError(err) {
	console.log(err.toString());
	this.emit('end');
}

gulp.task('compass', function() {
	gulp.src(paths.styles)
		.pipe(compass({
			config_file: './config.rb',
			sass: 'styles',
			css: paths.dist
		}))
		.on('error', handleError)
		.pipe(gulp.dest(paths.dist));
});

gulp.task('uglify', function() {
	gulp.src(paths.js)
		.pipe(uglify())
		.pipe(gulp.dest(paths.dist));
});

gulp.task('concat', function() {
	gulp.src(paths.libjs)
		.pipe(concat('app.js'))
		.pipe(gulp.dest(paths.dist));
});

gulp.task('htmlmin', function() {
	gulp.src(paths.htmls)
		.pipe(
			htmlmin({
				removeComments: true,
				removeCommentsFromCDATA: true,
				removeCDATASectionsFromCDATA: true,
				collapseWhitespace: true,
				collapseBooleanAttributes: true,
				removeAttributeQuotes: true,
				removeRedundantAttributes: true,
				removeEmptyAttributes: true,
				removeScriptTypeAttributes: true,
				removeOptionalTags: true
			})
		)
		.pipe(gulp.dest(paths.dist));
});

gulp.task('watch', function() {
	gulp.watch(paths.styles, ['compass']);
	gulp.watch(paths.js, ['uglify', 'concat']);
	gulp.watch(paths.htmls, ['htmlmin']);
});

gulp.task('default', ['compass', 'uglify', 'concat', 'htmlmin', 'watch']);

gulp.task('build', ['compass', 'uglify', 'concat', 'htmlmin']);
