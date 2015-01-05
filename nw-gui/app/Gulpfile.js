var gulp = require('gulp'),
	compass = require('gulp-compass'),
	htmlmin = require('gulp-htmlmin'),
	uglify = require('gulp-uglify');

var paths = {
	scripts: 'scripts/*.js',
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
		.pipe(gulp.dest(paths.dist))
		.pipe(connect.reload());
});

gulp.task('uglify', function() {
  gulp.src(paths.scripts)
    .pipe(uglify())
    .pipe(gulp.dest(paths.dist))
    .pipe(connect.reload());
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
	.pipe(gulp.dest(paths.dist))
	.pipe(connect.reload());
});

gulp.task('watch', function() {
	gulp.watch(paths.styles, ['compass']);
	gulp.watch(paths.scripts, ['uglify']);
	gulp.watch(paths.htmls, ['htmlmin']);
});

gulp.task('default', ['compass', 'uglify', 'htmlmin', 'watch']);