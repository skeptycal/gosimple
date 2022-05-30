package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/skeptycal/gosimple/cli/errorlogger"
)

var (
	log = errorlogger.New()
	pwd string

	// TODO: fix this
	fakePath = `/Users/skeptycal/go/src/github.com/skeptycal/gosimple/templates/repofiles`
	outfile  = filepath.Join(fakePath, "template_gitignore")
)

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Debug(err)
	}

}

// local path: Users/skeptycal/go/src/github.com/skeptycal/gosimple/templates/repofiles
// local path: $GOPATH/src/github.com/skeptycal/gosimple/templates/repofiles

func main() {

	gi := template.New("gitignore")

	t := githubTemplate

	minParse := "min: {{.Header}}"

	// fmt.Println(t)
	fmt.Println()

	gi.Parse(minParse) // TODO: this works ...
	gi.Execute(os.Stdout, t)

	gi.Parse(giFile) // TODO: but this doesn't ??? ... ok, so ... use Uppercase for field names =)
	gi.Execute(os.Stdout, t)

	// gi := template.Must(template.New(githubTemplate), nil)

	// WriteTemplateFile(outfile, gi, nil, gofile.NormalMode)

}

// WriteTemplateFile writes the given template to the
// named file, creating it if necessary. If the file
// does not exist, WriteTemplateFile creates it with
// permissions perm (before umask); otherwise
// WriteTemplateFile truncates it before writing,
// without changing permissions.
func WriteTemplateFile(name string, t *template.Template, data any, perm os.FileMode) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}

	err = t.Execute(f, data)

	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}

const (
	// use pwd instead ...
	repoFilesPath = "../../../repofiles"
)

// https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django,yarn,git

// func template.ParseFiles(filenames ...string) (*template.Template, error)

func debugPaths() {
	fmt.Println("pwd: ", pwd)
	fmt.Println("fake repo path: ", fakePath)
	fmt.Println("output file: ", outfile)
	fmt.Println()
}

type Book struct {
	Title     string
	Publisher string
	Year      int
}

func bookExample() {
	gi := template.New("gitignore")
	gi.Parse("External variable has the value [{{.}}]\n")
	gi.Execute(os.Stdout, "Amazing")
	b := Book{"The CSound Book", "MIT Press", 2002}
	gi.Execute(os.Stdout, b)
}

var (
	giFile = "{{.header}}\n\n{{.divider}}\n\n{{.personalItems}}\n\n{{.repoSpecificItems}}\n\n{{.securityItems}}\n\n{{.divider}}\n\n{{.giFile}}"

	githubTemplate = struct {
		Header            string
		divider           string
		personalItems     string
		repoSpecificItems string
		securityItems     string
		giFile            string
		// t                 string
	}{
		Header:            header,
		divider:           divider,
		personalItems:     personalItems,
		repoSpecificItems: repoSpecificItems,
		securityItems:     securityItems,
		giFile:            "", // giSample,
	}

	header = `# Copyright (c) 2021 Michael Treanor
# https://github.com/skeptycal
# MIT License

# Template Test
`

	personalItems = `### Personal ###
**/[Bb]ak/
**/*.bak
**/*temp
**/*tmp
**/.waka*
.vscode`

	repoSpecificItems = `### Repo Specific ###
**/idea.md`

	securityItems = `### Security ###
**/*[Tt]oken*
**/*[Pp]rivate*
**/*[Ss]ecret*
*history*
*hst*`

	divider = `############################################`

	giSample = `################################ Created by https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django,yarn,git
# Edit at https://www.toptal.com/developers/gitignore?templates=macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django,yarn,git

### Django ###
*.log
*.pot
*.pyc
**/__pycache__/
local_settings.py
db.sqlite3
db.sqlite3-journal
media

# If your build process includes running collectstatic, then you probably don't need or want to include staticfiles/
# in your Git repository. Update and uncomment the following line accordingly.
# <django-project-name>/staticfiles/

### Django.Python Stack ###
# Byte-compiled / optimized / DLL files
*.py[cod]
*$py.class

# C extensions
*.so

# Distribution / packaging
.Python
**/build/
develop-eggs/
**/dist/
**/downloads/
**/eggs/
.eggs/
**/parts/
**/sdist/
**/var/
**/wheels/
pip-wheel-metadata/
**/share/python-wheels/
*.egg-info/
.installed.cfg
*.egg
MANIFEST

# PyInstaller
#  Usually these files are written by a python script from a template
#  before PyInstaller builds the exe, so as to inject date/other infos into it.
*.manifest
*.spec

# Installer logs
pip-log.txt
pip-delete-this-directory.txt

# Unit test / coverage reports
**/htmlcov/
.tox/
.nox/
.coverage
.coverage.*
.cache
nosetests.xml
coverage.xml
*.cover
*.py,cover
.hypothesis/
.pytest_cache/
pytestdebug.log

# Translations
*.mo

# Django stuff:

# Flask stuff:
**/instance/
.webassets-cache

# Scrapy stuff:
.scrapy

# Sphinx documentation
**/docs/_build/
**/doc/_build/

# PyBuilder
**/target/

# Jupyter Notebook
.ipynb_checkpoints

# IPython
**/profile_default/
ipython_config.py

# pyenv
.python-version

# pipenv
#   According to pypa/pipenv#598, it is recommended to include Pipfile.lock in version control.
#   However, in case of collaboration, if having platform-specific dependencies or dependencies
#   having no cross-platform support, pipenv may install dependencies that don't work, or not
#   install all needed dependencies.
#Pipfile.lock

# poetry
#poetry.lock

# PEP 582; used by e.g. github.com/David-OConnor/pyflow
**/__pypackages__/

# Celery stuff
celerybeat-schedule
celerybeat.pid

# SageMath parsed files
*.sage.py

# Environments
# .env
.env/
.venv/
**/env/
**/venv/
**/ENV/
env.bak/
venv.bak/
pythonenv*

# Spyder project settings
.spyderproject
.spyproject

# Rope project settings
.ropeproject

# mkdocs documentation
/site

# mypy
.mypy_cache/
.dmypy.json
dmypy.json

# Pyre type checker
.pyre/

# pytype static type analyzer
.pytype/

# operating system-related files
# file properties cache/storage on macOS
*.DS_Store
# thumbnail cache on Windows
Thumbs.db

# profiling data
.prof


### Git ###
# Created by git for backups. To disable backups in Git:
# $ git config --global mergetool.keepBackup false
*.orig

# Created by git when using merge tools for conflicts
*.BACKUP.*
*.BASE.*
*.LOCAL.*
*.REMOTE.*
*_BACKUP_*.txt
*_BASE_*.txt
*_LOCAL_*.txt
*_REMOTE_*.txt

### Go ###
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.dylib

# Test binary, built with go test -c
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

### Go Patch ###
/vendor/
/Godeps/

### Linux ###
*~

# temporary files which can be created if a process still has a handle open of a deleted file
.fuse_hidden*

# KDE directory preferences
.directory

# Linux trash folder which might appear on any partition or disk
.Trash-*

# .nfs files are created when an open file is removed but is still being accessed
.nfs*

### macOS ###
# General
.DS_Store
.AppleDouble
.LSOverride

# Icon must end with two \r
Icon


# Thumbnails
._*

# Files that might appear in the root of a volume
.DocumentRevisions-V100
.fseventsd
.Spotlight-V100
.TemporaryItems
.Trashes
.VolumeIcon.icns
.com.apple.timemachine.donotpresent

# Directories potentially created on remote AFP share
.AppleDB
.AppleDesktop
Network Trash Folder
Temporary Items
.apdisk

### Node ###
# Logs
logs
npm-debug.log*
yarn-debug.log*
yarn-error.log*
lerna-debug.log*

# Diagnostic reports (https://nodejs.org/api/report.html)
report.[0-9]*.[0-9]*.[0-9]*.[0-9]*.json

# Runtime data
pids
*.pid
*.seed
*.pid.lock

# Directory for instrumented libs generated by jscoverage/JSCover
lib-cov

# Coverage directory used by tools like istanbul
coverage
*.lcov

# nyc test coverage
.nyc_output

# Grunt intermediate storage (https://gruntjs.com/creating-plugins#storing-task-files)
.grunt

# Bower dependency directory (https://bower.io/)
bower_components

# node-waf configuration
.lock-wscript

# Compiled binary addons (https://nodejs.org/api/addons.html)
**/build/Release

# Dependency directories
**/node_modules/
**/jspm_packages/

# TypeScript v1 declaration files
**/typings/

# TypeScript cache
*.tsbuildinfo

# Optional npm cache directory
.npm

# Optional eslint cache
.eslintcache

# Optional stylelint cache
.stylelintcache

# Microbundle cache
.rpt2_cache/
.rts2_cache_cjs/
.rts2_cache_es/
.rts2_cache_umd/

# Optional REPL history
.node_repl_history

# Output of 'npm pack'
*.tgz

# Yarn Integrity file
.yarn-integrity

# dotenv environment variables file
.env
.env.test
.env*.local

# parcel-bundler cache (https://parceljs.org/)
.parcel-cache

# Next.js build output
.next

# Nuxt.js build / generate output
.nuxt
dist

# Storybook build outputs
.out
.storybook-out
storybook-static

# rollup.js default build output

# Gatsby files
.cache/
# Comment in the public line in if your project uses Gatsby and not Next.js
# https://nextjs.org/blog/next-9-1#public-directory-support
# public

# vuepress build output
.vuepress/dist

# Serverless directories
.serverless/

# FuseBox cache
.fusebox/

# DynamoDB Local files
.dynamodb/

# TernJS port file
.tern-port

# Stores VSCode versions used for testing VSCode extensions
.vscode-test

# Temporary folders
**/tmp/
**/temp/

### Nuxt ###
# gitignore template for Nuxt.js projects
#
# Recommended template: Node.gitignore

# Nuxt build

# Nuxt generate

### Python ###
# Byte-compiled / optimized / DLL files

# C extensions

# Distribution / packaging

# PyInstaller
#  Usually these files are written by a python script from a template
#  before PyInstaller builds the exe, so as to inject date/other infos into it.

# Installer logs

# Unit test / coverage reports

# Translations

# Django stuff:

# Flask stuff:

# Scrapy stuff:

# Sphinx documentation

# PyBuilder

# Jupyter Notebook

# IPython

# pyenv

# pipenv
#   According to pypa/pipenv#598, it is recommended to include Pipfile.lock in version control.
#   However, in case of collaboration, if having platform-specific dependencies or dependencies
#   having no cross-platform support, pipenv may install dependencies that don't work, or not
#   install all needed dependencies.

# poetry

# PEP 582; used by e.g. github.com/David-OConnor/pyflow

# Celery stuff

# SageMath parsed files

# Environments
# .env

# Spyder project settings

# Rope project settings

# mkdocs documentation

# mypy

# Pyre type checker

# pytype static type analyzer

# operating system-related files
# file properties cache/storage on macOS
# thumbnail cache on Windows

# profiling data


### SSH ###
**/.ssh/
**/.ssh/id_*
**/.ssh/*_id_*
**/.ssh/known_hosts

### vscode ###
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
*.code-workspace

### Vue ###
# gitignore template for Vue.js projects
# Recommended template: Node.gitignore

# TODO: where does this rule come from?
**/docs/_book

# TODO: where does this rule come from?
**/test/

### Windows ###
# Windows thumbnail cache files
Thumbs.db:encryptable
ehthumbs.db
ehthumbs_vista.db

# Dump file
*.stackdump

# Folder config file
[Dd]esktop.ini

# Recycle Bin used on file shares
$RECYCLE.BIN/

# Windows Installer files
*.cab
*.msi
*.msix
*.msm
*.msp

# Windows shortcuts
*.lnk

### yarn ###
# https://yarnpkg.com/advanced/qa#which-files-should-be-gitignored

.yarn/*
!.yarn/releases
!.yarn/plugins
!.yarn/sdks
!.yarn/versions

# if you are NOT using Zero-installs, then:
# comment the following lines
!.yarn/cache

# and uncomment the following lines
# .pnp.*

### Zsh ###
# Zsh compiled script + zrecompile backup
*.zwc
*.zwc.old

# Zsh completion-optimization dumpfile
*zcompdump*

# Zsh zcalc history
.zcalc_history

# A popular plugin manager's files
._zinit
.zinit_lstupd

# zdharma/zshelldoc tool's files
**/zsdoc/data

# robbyrussell/oh-my-zsh/plugins/per-directory-history plugin's files
# (when set-up to store the history in the local directory)
.directory_history

# MichaelAquilina/zsh-autoswitch-virtualenv plugin's files
# (for Zsh plugins using Python)
.venv

# Zunit tests' output
/tests/_output/*
!/tests/_output/.gitkeep

# End of https://www.toptal.com/developers/gitignore/api/macos,linux,windows,ssh,vscode,go,zsh,node,vue,nuxt,python,django,yarn,git
`
)
