# gorepotemplate

> This is the initial directory tree for gorepotemplate. Use the make_tree_md.sh script ([GNU-tree required][get_tree]) to update it if you wish. It is safe to delete this file.

### Directory Structure

```sh
.
├── .editorconfig
├── .github
│   ├── FUNDING.yml
│   ├── ISSUE_TEMPLATE
│   │   ├── bug_report.md
│   │   └── feature_request.md
│   └── workflows
│       └── go.yml
├── .gitignore
├── .vscode
│   └── settings.json
├── CODE_OF_CONDUCT.md
├── LICENSE
├── README.md
├── SECURITY.md
├── cli
│   ├── README.md
│   ├── ansi
│   │   ├── .editorconfig
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── ansi_code.go
│   │   ├── ansi_code_test.go
│   │   ├── ansi_colors.md
│   │   ├── ascii.md
│   │   ├── assets
│   │   │   ├── ansi_backgrounds.jpg
│   │   │   └── ansi_foregrounds.jpg
│   │   ├── benchmark
│   │   │   ├── benchmark.go
│   │   │   ├── benchmarkset.go
│   │   │   ├── config.go
│   │   │   ├── httpresponse.go
│   │   │   ├── testdata.go
│   │   │   ├── testfuncs.go
│   │   │   └── verbose.go
│   │   ├── cmd
│   │   │   ├── ansistrings
│   │   │   │   └── main.go
│   │   │   ├── bytes
│   │   │   │   ├── bytes
│   │   │   │   ├── main
│   │   │   │   └── main.go
│   │   │   └── gofileparse
│   │   │       └── main.go
│   │   ├── constants.go
│   │   ├── contributing.md
│   │   ├── example.go
│   │   ├── go.test.sh
│   │   ├── join
│   │   │   ├── join.go
│   │   │   └── join_test.go
│   │   ├── sample_benchmarks
│   │   │   └── fib
│   │   │       ├── fib
│   │   │       │   ├── fib.go
│   │   │       │   └── fib_test.go
│   │   │       └── main.go
│   │   └── symbols.go
│   ├── ansi.go
│   ├── ansilogger
│   │   ├── ansi.go
│   │   └── ansi_test.go
│   ├── cli
│   │   └── main.go
│   ├── cli.go
│   ├── cli_controls.go
│   ├── constants.go
│   ├── constants_test.go
│   ├── cursor.go
│   ├── encode.go
│   ├── encode_test.go
│   ├── envvars
│   │   └── envvars.go
│   ├── errorlogger
│   │   ├── .VERSION
│   │   ├── .editorconfig
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── docs
│   │   │   ├── _config.yml
│   │   │   ├── docs.md
│   │   │   ├── index.html
│   │   │   └── template.md
│   │   ├── error_func.go
│   │   ├── error_func_test.go
│   │   ├── errorlogger.go
│   │   ├── errorlogger_test.go
│   │   ├── example.go
│   │   ├── go.test.sh
│   │   ├── idea.md
│   │   ├── internal.go
│   │   ├── json_formatter.go
│   │   ├── json_formatter_test.go
│   │   ├── level.go
│   │   ├── logrus_types.go
│   │   ├── test_info_test.go
│   │   ├── text_formatter.go
│   │   └── text_formatter_test.go
│   ├── goshell
│   │   ├── .editorconfig
│   │   ├── .pre-commit-config.yaml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── cmd
│   │   │   └── example
│   │   │       └── goshell
│   │   │           └── main.go
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── defaults.go
│   │   ├── defaults_test.go
│   │   ├── docs
│   │   │   ├── _config.yml
│   │   │   ├── docs.md
│   │   │   ├── index.html
│   │   │   └── template.md
│   │   ├── env.go
│   │   ├── example.go
│   │   ├── go.test.sh
│   │   ├── homedir.go
│   │   ├── homedir_test.go
│   │   ├── idea.md
│   │   ├── internal
│   │   │   └── fixture
│   │   │       └── test.go
│   │   ├── make_tree_md.sh
│   │   ├── profile1651526871727910000.out
│   │   ├── shell.go
│   │   ├── shell_test.go
│   │   └── tree.md
│   ├── internal.go
│   ├── internal_test.go
│   ├── miniansi
│   │   ├── .editorconfig
│   │   ├── .vscode
│   │   │   └── settings.json
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── constants.go
│   │   ├── contributing.md
│   │   ├── example.go
│   │   ├── miniansi.go
│   │   ├── miniansi_test.go
│   │   └── types.go
│   ├── osargs
│   │   ├── .editorconfig
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── example.go
│   │   ├── go.test.sh
│   │   ├── osargs.go
│   │   └── osargs_test.go
│   ├── printer
│   │   ├── printer.go
│   │   └── printer_test.go
│   ├── term.go
│   ├── terminal
│   │   ├── terminal.go
│   │   ├── terminal_check_appengine.go
│   │   ├── terminal_check_bsd.go
│   │   ├── terminal_check_js.go
│   │   ├── terminal_check_no_terminal.go
│   │   ├── terminal_check_notappengine.go
│   │   ├── terminal_check_solaris.go
│   │   ├── terminal_check_unix.go
│   │   ├── terminal_check_windows.go
│   │   ├── terminal_nosysioctl.go
│   │   ├── terminal_sysioctl.go
│   │   └── terminal_windows.go
│   ├── terminal.go
│   ├── util.go
│   └── util_test.go
├── cmd
│   ├── examples
│   │   ├── delsubs
│   │   │   ├── find.help.out
│   │   │   ├── find.helpD.man
│   │   │   ├── find.man
│   │   │   └── main.go
│   │   ├── errorlogger
│   │   │   └── main.go
│   │   ├── miniansi
│   │   │   └── main.go
│   │   └── osargs
│   │       └── main.go
│   ├── googleapiquickstart
│   │   ├── .VERSION
│   │   ├── .gitignore
│   │   ├── client.go
│   │   ├── cmd
│   │   │   └── example
│   │   │       └── main.go
│   │   └── example.go
│   └── goutil_playground
│       ├── asserts.go
│       ├── bubblesort_test.go
│       ├── dict
│       │   ├── components.go
│       │   ├── dict.go
│       │   ├── dict_test.go
│       │   ├── example
│       │   │   └── main.go
│       │   └── examples
│       │       ├── addone
│       │       │   ├── examples
│       │       │   │   └── examples.go
│       │       │   ├── main.go
│       │       │   └── main_test.go
│       │       └── genericfunc
│       │           └── main.go
│       ├── generic.go
│       ├── generic_test.go
│       ├── kinds.go
│       ├── list
│       │   ├── bubblesort.go
│       │   └── list.go
│       ├── manage
│       │   ├── .testoutput
│       │   │   └── heap.txt
│       │   ├── data
│       │   │   ├── aliases.txt
│       │   │   ├── coverage.txt
│       │   │   ├── env.txt
│       │   │   └── gh_repo_create_help.txt
│       │   ├── examples
│       │   │   ├── escapes
│       │   │   │   └── main.go
│       │   │   └── gitit
│       │   │       ├── main.go
│       │   │       ├── main_test.go
│       │   │       └── temp
│       │   ├── ghshell
│       │   │   ├── .testoutput
│       │   │   │   └── heap.txt
│       │   │   ├── cmd.go
│       │   │   ├── errorcontrol.go
│       │   │   ├── errors.go
│       │   │   ├── ghshell.go
│       │   │   ├── ghshell_test.go
│       │   │   ├── gitit.go
│       │   │   ├── gobot.go
│       │   │   ├── initgobot.go
│       │   │   ├── logging.go
│       │   │   ├── rand.go
│       │   │   ├── unix_rusage.1
│       │   │   ├── util.go
│       │   │   └── util_test.go
│       │   ├── gnuflags
│       │   │   └── flags.go
│       │   ├── main.go
│       │   ├── profile.out
│       │   └── shellscripts
│       │       ├── gitit.sh
│       │       ├── gitsub.sh
│       │       ├── gitutil.sh
│       │       ├── go.test.sh
│       │       ├── gobuildflagsoutput.sh
│       │       ├── update.sh
│       │       ├── updatesubs.sh
│       │       └── workspace_init.sh
│       ├── nongeneric
│       │   └── nongeneric.go
│       ├── sampletypes.go
│       ├── sequence
│       │   ├── examples.go
│       │   ├── sequence.go
│       │   └── sequence_test.go
│       ├── so_answer1.md
│       ├── so_questions.go
│       ├── soexamples
│       │   ├── 71677581
│       │   │   ├── .editorconfig
│       │   │   ├── CODE_OF_CONDUCT.md
│       │   │   ├── LICENSE
│       │   │   ├── README.md
│       │   │   ├── SECURITY.md
│       │   │   ├── contributing.md
│       │   │   ├── coverage.txt
│       │   │   ├── go.test.sh
│       │   │   ├── main.go
│       │   │   ├── main.go.bak
│       │   │   ├── main_test.go
│       │   │   └── profile.out
│       │   ├── othermain.go
│       │   └── soexamples.bak
│       │       ├── main.go
│       │       ├── main.go.bak
│       │       └── main_test.go
│       ├── sort
│       │   ├── floats.go
│       │   ├── generic.go
│       │   ├── ints.go
│       │   ├── sort.go
│       │   ├── stable.go
│       │   ├── stdlib.go
│       │   ├── stdlibinternal.go
│       │   └── strings.go
│       ├── stack.go
│       ├── stack_test.go
│       └── types.go
├── contributing.md
├── coverage.txt
├── datatools
│   ├── atomic
│   │   ├── atomic.go
│   │   ├── atomic_test.go
│   │   ├── endian_big.go
│   │   └── endian_little.go
│   ├── binary
│   │   └── percentages.go
│   ├── bufferpool
│   │   ├── bufferpool.go
│   │   ├── bufferpool_test.go
│   │   ├── coverage.txt
│   │   ├── example
│   │   │   ├── main.go
│   │   │   └── main_test.go
│   │   ├── go.doc
│   │   ├── pool.go
│   │   ├── poolreset.go
│   │   ├── profile1651586595841212000.out
│   │   └── syncpool_test.go
│   ├── cmd
│   │   └── EmailDomains
│   │       ├── emaildomains
│   │       └── main.go
│   ├── compare
│   │   └── interface.go
│   ├── data
│   │   ├── cpuoptions.go
│   │   ├── database.go
│   │   ├── romeo_and_juliet.txt
│   │   └── sql.go
│   ├── format
│   │   ├── email.go
│   │   ├── email_test.go
│   │   ├── numberformatting.go
│   │   ├── numberformatting_test.go
│   │   └── sample.txt
│   ├── list
│   │   ├── builder.go
│   │   ├── builder_test.go
│   │   ├── insert.go
│   │   ├── intbuilder_test.go
│   │   ├── internal.go
│   │   ├── list.go
│   │   ├── list_test.go
│   │   └── remove.go
│   ├── math
│   │   └── polynomial
│   │       ├── cmd
│   │       │   └── poly
│   │       │       ├── main.go
│   │       │       └── poly
│   │       ├── polynomial.go
│   │       └── polynomial_test.go
│   ├── mysql
│   │   ├── LICENSE
│   │   ├── config.go
│   │   ├── license.go
│   │   └── mysql.go
│   ├── primes10k
│   │   └── primes10k.go
│   ├── rand
│   │   └── rand.go
│   ├── sequence
│   │   ├── sequence
│   │   └── sequence_test.go
│   ├── sieve
│   │   ├── .vscode
│   │   │   └── settings.json
│   │   ├── examples
│   │   │   ├── basic
│   │   │   │   └── main.go
│   │   │   ├── channel_flush
│   │   │   │   └── main.go
│   │   │   ├── channel_flush2
│   │   │   │   └── main.go
│   │   │   ├── channel_flush3
│   │   │   │   └── main.go
│   │   │   ├── channel_flush4
│   │   │   │   └── main.go
│   │   │   ├── channels
│   │   │   │   ├── channels.go
│   │   │   │   ├── channels_test.go
│   │   │   │   ├── heap.0.pprof
│   │   │   │   ├── heap.1.pprof
│   │   │   │   ├── heap.2.pprof
│   │   │   │   ├── heap.3.pprof
│   │   │   │   ├── heap.4.pprof
│   │   │   │   ├── pprof.sh
│   │   │   │   └── reconnect.go
│   │   │   └── counting
│   │   │       └── main.go
│   │   ├── primetest.go
│   │   ├── primetest_test.go
│   │   ├── sieve.go
│   │   ├── sieveboolmap.go
│   │   └── sievecounting.go
│   └── size
│       ├── size.go
│       └── size_test.go
├── docs
│   ├── _config.yml
│   ├── docs.md
│   ├── index.html
│   └── template.md
├── go.mod
├── go.sum
├── go.test.sh
├── gogit
│   ├── _example
│   │   ├── cli_examples
│   │   │   └── github_api_response.json
│   │   └── escape-seq
│   │       └── main.go
│   ├── gogit.go
│   └── gogit_test.go
├── io
│   ├── examples
│   │   └── comments
│   │       └── main.go
│   ├── reader.go
│   ├── scanner.go
│   └── writer.go
├── make_tree_md.sh
├── os
│   ├── basicfile
│   │   ├── basicfile.go
│   │   ├── datafile.go
│   │   ├── direntry.go
│   │   ├── errors.go
│   │   ├── fileinfo.go
│   │   ├── fileinfoaliases.go
│   │   ├── filemode.go
│   │   ├── fileops.go
│   │   ├── fileunix.go
│   │   ├── gofile.go
│   │   ├── gofileerror.go
│   │   ├── internal.go
│   │   ├── llrb_avg.go
│   │   ├── llrb_avg_test.go
│   │   ├── textfile.go
│   │   ├── types.go
│   │   ├── util.go
│   │   └── ~bench_results.csv
│   ├── gofile
│   │   ├── .editorconfig
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── VERSION
│   │   ├── cmd
│   │   ├── constants.go
│   │   ├── contributing.md
│   │   ├── copy.go
│   │   ├── copybenchmarks
│   │   │   ├── copy_test.go
│   │   │   ├── fakeDst
│   │   │   └── fakeSrc
│   │   ├── coverage.txt
│   │   ├── dir_options.go
│   │   ├── dirlist.go
│   │   ├── errors.go
│   │   ├── fileops.go
│   │   ├── go.test.sh
│   │   ├── gofile.go
│   │   ├── gofileerror.go
│   │   ├── gofileerror_test.go
│   │   ├── internal.go
│   │   ├── internal_test.go
│   │   ├── logging.go
│   │   ├── profile.out
│   │   └── types.go
│   ├── ls
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── ansi.go
│   │   ├── contributing.md
│   │   ├── git.go
│   │   ├── ls.go
│   │   ├── ls.txt
│   │   └── main.go
│   ├── memfile
│   │   ├── LICENSE
│   │   ├── memfile.go
│   │   ├── memfile_test.go
│   │   ├── uniuri.go
│   │   ├── util.go
│   │   └── util_test.go
│   ├── redlogger
│   │   └── redlogger.go
│   └── shellpath
│       ├── cmd
│       │   └── path.go
│       ├── osutil.go
│       ├── shpath.go
│       ├── shpath2.go
│       ├── shpath_test.go
│       └── stringutil.go
├── profile1651170016096182000.out
├── profile1651526732102151000.out
├── profile1651860465558384000.out
├── repo
│   ├── defaults
│   │   ├── .editorconfig
│   │   ├── .pre-commit-config.yaml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── ansi.go
│   │   ├── cmd
│   │   │   ├── benchmarks
│   │   │   │   └── formatbenchmarks
│   │   │   │       ├── main.go
│   │   │   │       └── main_test.go
│   │   │   └── example
│   │   │       ├── color_example
│   │   │       │   └── main.go
│   │   │       └── defaults
│   │   │           └── main.go
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── debug.go
│   │   ├── debug_test.go
│   │   ├── defaults.go
│   │   ├── defaults_test.go
│   │   ├── docs
│   │   │   ├── _config.yml
│   │   │   ├── docs.md
│   │   │   ├── index.html
│   │   │   └── template.md
│   │   ├── example.go
│   │   ├── flags.go
│   │   ├── go.test.sh
│   │   ├── idea.md
│   │   ├── internal.go
│   │   ├── make_tree_md.sh
│   │   ├── settings.go
│   │   ├── test_utils.go
│   │   ├── test_utils_test.go
│   │   ├── trace.go
│   │   ├── tree.md
│   │   ├── types.go
│   │   ├── types_test.go
│   │   ├── utils.go
│   │   └── utils_test.go
│   ├── fake_repo
│   │   ├── LICENSE
│   │   └── fakerepo.go
│   ├── fakecloser
│   │   └── fakecloser.go
│   ├── ghrepo
│   ├── gitcommits
│   │   ├── .editorconfig
│   │   ├── .pre-commit-config.yaml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── cmd
│   │   │   └── example
│   │   │       └── gitcommits
│   │   │           └── gitcommits.go
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── docs
│   │   │   ├── _config.yml
│   │   │   ├── docs.md
│   │   │   ├── index.html
│   │   │   └── template.md
│   │   ├── example.go
│   │   ├── gitcommits.go
│   │   ├── go.test.sh
│   │   ├── idea.md
│   │   ├── make_tree_md.sh
│   │   └── tree.md
│   ├── gitignore
│   │   ├── gi_fileparser
│   │   │   ├── gilist.txt
│   │   │   └── main.go
│   │   ├── gitignore.go
│   │   └── gitignore_gen.go
│   ├── gitroot
│   │   ├── find
│   │   │   ├── find.go
│   │   │   └── find_test.go
│   │   └── root.go
│   ├── go-version
│   │   ├── .circleci
│   │   │   └── config.yml
│   │   ├── CHANGELOG.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── constraint.go
│   │   ├── constraint_test.go
│   │   ├── version.go
│   │   ├── version_collection.go
│   │   ├── version_collection_test.go
│   │   └── version_test.go
│   ├── goconfig
│   │   ├── flags.go
│   │   ├── goconfig.go
│   │   └── goconfig_test.go
│   ├── gogithub
│   │   ├── .editorconfig
│   │   ├── .pre-commit-config.yaml
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── README.md
│   │   ├── SECURITY.md
│   │   ├── auth.go
│   │   ├── builderpool.go
│   │   ├── cmd
│   │   │   └── example
│   │   │       └── gogithub
│   │   │           └── example.go
│   │   ├── contributing.md
│   │   ├── coverage.txt
│   │   ├── docs
│   │   │   ├── _config.yml
│   │   │   ├── docs.md
│   │   │   ├── index.html
│   │   │   └── template.md
│   │   ├── example.go
│   │   ├── exec.go
│   │   ├── exec_test.go
│   │   ├── go.test.sh
│   │   ├── gogithub.go
│   │   ├── idea.md
│   │   ├── job.go
│   │   ├── job_test.go
│   │   ├── make_tree_md.sh
│   │   ├── sbWriter.go
│   │   ├── sbpool.go
│   │   └── tree.md
│   ├── gomake
│   │   ├── datafile.go
│   │   ├── example.go
│   │   ├── gi_list
│   │   ├── gogithub.go
│   │   ├── gomake.go
│   │   ├── internal.go
│   │   ├── pprint.go
│   │   ├── template_files
│   │   │   ├── README.md
│   │   │   ├── index.html
│   │   │   └── screenshot.css
│   │   ├── templates.go
│   │   └── util.go
│   ├── gomod
│   │   ├── .gitmodules
│   │   ├── checks
│   │   │   └── checks.go
│   │   ├── gen
│   │   │   └── genzfunc.go
│   │   ├── main.go
│   │   └── mod
│   │       ├── mod.go
│   │       └── mod_test.go
│   ├── repo_management
│   │   ├── cleanlist
│   │   │   └── cleanlist.go
│   │   ├── config
│   │   │   ├── config.go
│   │   │   ├── flags.go
│   │   │   ├── syncmap.go
│   │   │   └── util.go
│   │   ├── countargs.sh
│   │   ├── del_list.sh
│   │   ├── deleted_4_16_22.ini
│   │   ├── delrepos.sh
│   │   ├── forks.csv
│   │   ├── forks_list.csv
│   │   ├── ghnamecheck
│   │   │   ├── ghnamecheck_test.go
│   │   │   └── main.go
│   │   ├── make_repo_list.sh
│   │   ├── sources.csv
│   │   └── sources_list.csv
│   ├── seeker
│   │   ├── cli.go
│   │   ├── cmd
│   │   │   └── server
│   │   │       ├── serve.go
│   │   │       └── serve_test.go
│   │   ├── config.go
│   │   ├── handlers.go
│   │   ├── http.go
│   │   ├── mapper.go
│   │   ├── seeker.go
│   │   ├── seeker.sh
│   │   ├── time.go
│   │   ├── timeMap.go
│   │   ├── timemapper.go
│   │   └── util.go
│   ├── template
│   │   ├── LICENSE
│   │   └── license.go
│   └── version
│       ├── sort.go
│       └── version.go
├── tests
│   ├── coverage.txt
│   ├── fields.go
│   ├── stringasserts.go
│   ├── stringasserts_test.go
│   ├── trun.go
│   └── value.go
├── tree.md
├── types
│   ├── .VERSION
│   ├── GOOSlist.txt
│   ├── any.go
│   ├── any_test.go
│   ├── anyvalue.go
│   ├── anyvalue_test.go
│   ├── channels
│   │   ├── concurrent
│   │   │   ├── dict.go
│   │   │   ├── kvstore.go
│   │   │   ├── list.go
│   │   │   ├── rand.go
│   │   │   ├── rand_test.go
│   │   │   └── util.go
│   │   └── main.go
│   ├── cmd
│   │   └── examples
│   │       ├── cpu
│   │       │   └── main.go
│   │       ├── ls
│   │       │   ├── ls
│   │       │   └── main.go
│   │       ├── terminal
│   │       │   └── main.go
│   │       ├── text_analysis
│   │       │   └── letter_frequency
│   │       │       ├── main.go
│   │       │       └── romeo_and_juliet.txt
│   │       └── values
│   │           └── main.go
│   ├── const.go
│   ├── const_test.go
│   ├── constraints
│   │   ├── CODE_OF_CONDUCT.md
│   │   ├── LICENSE
│   │   ├── SECURITY.md
│   │   ├── constraints.go
│   │   ├── constraints_test.go
│   │   ├── coverage.txt
│   │   ├── go.test.sh
│   │   └── helpers.go
│   ├── convert
│   │   ├── convert.go
│   │   ├── example
│   │   │   └── main.go
│   │   ├── numberer.go
│   │   ├── types.go
│   │   └── unsafe.go
│   ├── coverage.txt
│   ├── cpuoptions_test.go
│   ├── go.test.sh
│   ├── greek.go
│   ├── greek_test.go
│   ├── hook.go
│   ├── interfaces.go
│   ├── internal.go
│   ├── internal_test.go
│   ├── kindinfo.go
│   ├── parse.go
│   ├── reflections.go
│   ├── reflections_test.go
│   ├── temp.go
│   ├── temp_test.go
│   ├── terminal.go
│   ├── terminal_test.go
│   ├── test_assert.go
│   ├── test_utils.go
│   ├── test_utils_test.go
│   ├── tester.go
│   ├── testset.go
│   ├── text.go
│   ├── text_test.go
│   ├── types.go
│   ├── types_test.go
│   ├── util.go
│   ├── util_test.go
│   └── writer.go
├── vendor
│   ├── cloud.google.com
│   │   └── go
│   │       ├── LICENSE
│   │       └── compute
│   │           └── metadata
│   │               └── metadata.go
│   ├── github.com
│   │   ├── buger
│   │   │   └── goterm
│   │   │       ├── .gitignore
│   │   │       ├── LICENSE
│   │   │       ├── README.md
│   │   │       ├── box.go
│   │   │       ├── plot.go
│   │   │       ├── table.go
│   │   │       ├── terminal.go
│   │   │       ├── terminal_nosysioctl.go
│   │   │       ├── terminal_sysioctl.go
│   │   │       └── terminal_windows.go
│   │   ├── cpuguy83
│   │   │   └── go-md2man
│   │   │       └── v2
│   │   │           ├── LICENSE.md
│   │   │           └── md2man
│   │   │               ├── md2man.go
│   │   │               └── roff.go
│   │   ├── fatih
│   │   │   ├── color
│   │   │   │   ├── LICENSE.md
│   │   │   │   ├── README.md
│   │   │   │   ├── color.go
│   │   │   │   └── doc.go
│   │   │   └── structs
│   │   │       ├── .gitignore
│   │   │       ├── .travis.yml
│   │   │       ├── LICENSE
│   │   │       ├── README.md
│   │   │       ├── field.go
│   │   │       ├── structs.go
│   │   │       └── tags.go
│   │   ├── go-sql-driver
│   │   │   └── mysql
│   │   │       ├── .gitignore
│   │   │       ├── AUTHORS
│   │   │       ├── CHANGELOG.md
│   │   │       ├── LICENSE
│   │   │       ├── README.md
│   │   │       ├── auth.go
│   │   │       ├── buffer.go
│   │   │       ├── collations.go
│   │   │       ├── conncheck.go
│   │   │       ├── conncheck_dummy.go
│   │   │       ├── connection.go
│   │   │       ├── connector.go
│   │   │       ├── const.go
│   │   │       ├── driver.go
│   │   │       ├── dsn.go
│   │   │       ├── errors.go
│   │   │       ├── fields.go
│   │   │       ├── fuzz.go
│   │   │       ├── infile.go
│   │   │       ├── nulltime.go
│   │   │       ├── nulltime_go113.go
│   │   │       ├── nulltime_legacy.go
│   │   │       ├── packets.go
│   │   │       ├── result.go
│   │   │       ├── rows.go
│   │   │       ├── statement.go
│   │   │       ├── transaction.go
│   │   │       └── utils.go
│   │   ├── golang
│   │   │   ├── groupcache
│   │   │   │   ├── LICENSE
│   │   │   │   └── lru
│   │   │   │       └── lru.go
│   │   │   └── protobuf
│   │   │       ├── AUTHORS
│   │   │       ├── CONTRIBUTORS
│   │   │       ├── LICENSE
│   │   │       ├── proto
│   │   │       │   ├── buffer.go
│   │   │       │   ├── defaults.go
│   │   │       │   ├── deprecated.go
│   │   │       │   ├── discard.go
│   │   │       │   ├── extensions.go
│   │   │       │   ├── properties.go
│   │   │       │   ├── proto.go
│   │   │       │   ├── registry.go
│   │   │       │   ├── text_decode.go
│   │   │       │   ├── text_encode.go
│   │   │       │   ├── wire.go
│   │   │       │   └── wrappers.go
│   │   │       └── ptypes
│   │   │           ├── any
│   │   │           │   └── any.pb.go
│   │   │           ├── any.go
│   │   │           ├── doc.go
│   │   │           ├── duration
│   │   │           │   └── duration.pb.go
│   │   │           ├── duration.go
│   │   │           ├── timestamp
│   │   │           │   └── timestamp.pb.go
│   │   │           └── timestamp.go
│   │   ├── google
│   │   │   ├── go-github
│   │   │   │   └── v34
│   │   │   │       ├── AUTHORS
│   │   │   │       ├── LICENSE
│   │   │   │       └── github
│   │   │   │           ├── actions.go
│   │   │   │           ├── actions_artifacts.go
│   │   │   │           ├── actions_runner_groups.go
│   │   │   │           ├── actions_runners.go
│   │   │   │           ├── actions_secrets.go
│   │   │   │           ├── actions_workflow_jobs.go
│   │   │   │           ├── actions_workflow_runs.go
│   │   │   │           ├── actions_workflows.go
│   │   │   │           ├── activity.go
│   │   │   │           ├── activity_events.go
│   │   │   │           ├── activity_notifications.go
│   │   │   │           ├── activity_star.go
│   │   │   │           ├── activity_watching.go
│   │   │   │           ├── admin.go
│   │   │   │           ├── admin_orgs.go
│   │   │   │           ├── admin_stats.go
│   │   │   │           ├── admin_users.go
│   │   │   │           ├── apps.go
│   │   │   │           ├── apps_installation.go
│   │   │   │           ├── apps_manifest.go
│   │   │   │           ├── apps_marketplace.go
│   │   │   │           ├── authorizations.go
│   │   │   │           ├── billing.go
│   │   │   │           ├── checks.go
│   │   │   │           ├── code-scanning.go
│   │   │   │           ├── doc.go
│   │   │   │           ├── enterprise.go
│   │   │   │           ├── enterprise_actions_runners.go
│   │   │   │           ├── enterprise_audit_log.go
│   │   │   │           ├── event.go
│   │   │   │           ├── event_types.go
│   │   │   │           ├── gists.go
│   │   │   │           ├── gists_comments.go
│   │   │   │           ├── git.go
│   │   │   │           ├── git_blobs.go
│   │   │   │           ├── git_commits.go
│   │   │   │           ├── git_refs.go
│   │   │   │           ├── git_tags.go
│   │   │   │           ├── git_trees.go
│   │   │   │           ├── github-accessors.go
│   │   │   │           ├── github.go
│   │   │   │           ├── gitignore.go
│   │   │   │           ├── interactions.go
│   │   │   │           ├── interactions_orgs.go
│   │   │   │           ├── interactions_repos.go
│   │   │   │           ├── issue_import.go
│   │   │   │           ├── issues.go
│   │   │   │           ├── issues_assignees.go
│   │   │   │           ├── issues_comments.go
│   │   │   │           ├── issues_events.go
│   │   │   │           ├── issues_labels.go
│   │   │   │           ├── issues_milestones.go
│   │   │   │           ├── issues_timeline.go
│   │   │   │           ├── licenses.go
│   │   │   │           ├── messages.go
│   │   │   │           ├── migrations.go
│   │   │   │           ├── migrations_source_import.go
│   │   │   │           ├── migrations_user.go
│   │   │   │           ├── misc.go
│   │   │   │           ├── orgs.go
│   │   │   │           ├── orgs_actions_allowed.go
│   │   │   │           ├── orgs_actions_permissions.go
│   │   │   │           ├── orgs_audit_log.go
│   │   │   │           ├── orgs_hooks.go
│   │   │   │           ├── orgs_members.go
│   │   │   │           ├── orgs_outside_collaborators.go
│   │   │   │           ├── orgs_projects.go
│   │   │   │           ├── orgs_users_blocking.go
│   │   │   │           ├── packages.go
│   │   │   │           ├── projects.go
│   │   │   │           ├── pulls.go
│   │   │   │           ├── pulls_comments.go
│   │   │   │           ├── pulls_reviewers.go
│   │   │   │           ├── pulls_reviews.go
│   │   │   │           ├── reactions.go
│   │   │   │           ├── repos.go
│   │   │   │           ├── repos_collaborators.go
│   │   │   │           ├── repos_comments.go
│   │   │   │           ├── repos_commits.go
│   │   │   │           ├── repos_community_health.go
│   │   │   │           ├── repos_contents.go
│   │   │   │           ├── repos_deployments.go
│   │   │   │           ├── repos_forks.go
│   │   │   │           ├── repos_hooks.go
│   │   │   │           ├── repos_invitations.go
│   │   │   │           ├── repos_keys.go
│   │   │   │           ├── repos_merging.go
│   │   │   │           ├── repos_pages.go
│   │   │   │           ├── repos_prereceive_hooks.go
│   │   │   │           ├── repos_projects.go
│   │   │   │           ├── repos_releases.go
│   │   │   │           ├── repos_stats.go
│   │   │   │           ├── repos_statuses.go
│   │   │   │           ├── repos_traffic.go
│   │   │   │           ├── search.go
│   │   │   │           ├── strings.go
│   │   │   │           ├── teams.go
│   │   │   │           ├── teams_discussion_comments.go
│   │   │   │           ├── teams_discussions.go
│   │   │   │           ├── teams_members.go
│   │   │   │           ├── timestamp.go
│   │   │   │           ├── users.go
│   │   │   │           ├── users_administration.go
│   │   │   │           ├── users_blocking.go
│   │   │   │           ├── users_emails.go
│   │   │   │           ├── users_followers.go
│   │   │   │           ├── users_gpg_keys.go
│   │   │   │           ├── users_keys.go
│   │   │   │           ├── users_projects.go
│   │   │   │           ├── with_appengine.go
│   │   │   │           └── without_appengine.go
│   │   │   └── go-querystring
│   │   │       ├── LICENSE
│   │   │       └── query
│   │   │           └── encode.go
│   │   ├── googleapis
│   │   │   └── gax-go
│   │   │       └── v2
│   │   │           ├── LICENSE
│   │   │           ├── call_option.go
│   │   │           ├── gax.go
│   │   │           ├── header.go
│   │   │           └── invoke.go
│   │   ├── integralist
│   │   │   └── go-findroot
│   │   │       ├── LICENSE
│   │   │       └── find
│   │   │           └── find.go
│   │   ├── mattn
│   │   │   ├── go-colorable
│   │   │   │   ├── .travis.yml
│   │   │   │   ├── LICENSE
│   │   │   │   ├── README.md
│   │   │   │   ├── colorable_appengine.go
│   │   │   │   ├── colorable_others.go
│   │   │   │   ├── colorable_windows.go
│   │   │   │   ├── go.test.sh
│   │   │   │   └── noncolorable.go
│   │   │   └── go-isatty
│   │   │       ├── LICENSE
│   │   │       ├── README.md
│   │   │       ├── doc.go
│   │   │       ├── go.test.sh
│   │   │       ├── isatty_bsd.go
│   │   │       ├── isatty_others.go
│   │   │       ├── isatty_plan9.go
│   │   │       ├── isatty_solaris.go
│   │   │       ├── isatty_tcgets.go
│   │   │       └── isatty_windows.go
│   │   ├── mitchellh
│   │   │   └── go-homedir
│   │   │       ├── LICENSE
│   │   │       ├── README.md
│   │   │       └── homedir.go
│   │   ├── pkg
│   │   │   └── errors
│   │   │       ├── .gitignore
│   │   │       ├── .travis.yml
│   │   │       ├── LICENSE
│   │   │       ├── Makefile
│   │   │       ├── README.md
│   │   │       ├── appveyor.yml
│   │   │       ├── errors.go
│   │   │       ├── go113.go
│   │   │       └── stack.go
│   │   ├── russross
│   │   │   └── blackfriday
│   │   │       └── v2
│   │   │           ├── .gitignore
│   │   │           ├── .travis.yml
│   │   │           ├── LICENSE.txt
│   │   │           ├── README.md
│   │   │           ├── block.go
│   │   │           ├── doc.go
│   │   │           ├── esc.go
│   │   │           ├── html.go
│   │   │           ├── inline.go
│   │   │           ├── markdown.go
│   │   │           ├── node.go
│   │   │           └── smartypants.go
│   │   ├── shurcooL
│   │   │   ├── githubv4
│   │   │   │   ├── .travis.yml
│   │   │   │   ├── LICENSE
│   │   │   │   ├── README.md
│   │   │   │   ├── doc.go
│   │   │   │   ├── enum.go
│   │   │   │   ├── githubv4.go
│   │   │   │   ├── input.go
│   │   │   │   └── scalar.go
│   │   │   ├── graphql
│   │   │   │   ├── .travis.yml
│   │   │   │   ├── LICENSE
│   │   │   │   ├── README.md
│   │   │   │   ├── doc.go
│   │   │   │   ├── graphql.go
│   │   │   │   ├── ident
│   │   │   │   │   └── ident.go
│   │   │   │   ├── internal
│   │   │   │   │   └── jsonutil
│   │   │   │   │       └── graphql.go
│   │   │   │   ├── query.go
│   │   │   │   └── scalar.go
│   │   │   └── sanitized_anchor_name
│   │   │       ├── .travis.yml
│   │   │       ├── LICENSE
│   │   │       ├── README.md
│   │   │       └── main.go
│   │   ├── sirupsen
│   │   │   └── logrus
│   │   │       ├── .gitignore
│   │   │       ├── .golangci.yml
│   │   │       ├── .travis.yml
│   │   │       ├── CHANGELOG.md
│   │   │       ├── LICENSE
│   │   │       ├── README.md
│   │   │       ├── alt_exit.go
│   │   │       ├── appveyor.yml
│   │   │       ├── buffer_pool.go
│   │   │       ├── doc.go
│   │   │       ├── entry.go
│   │   │       ├── exported.go
│   │   │       ├── formatter.go
│   │   │       ├── hooks.go
│   │   │       ├── json_formatter.go
│   │   │       ├── logger.go
│   │   │       ├── logrus.go
│   │   │       ├── terminal_check_appengine.go
│   │   │       ├── terminal_check_bsd.go
│   │   │       ├── terminal_check_js.go
│   │   │       ├── terminal_check_no_terminal.go
│   │   │       ├── terminal_check_notappengine.go
│   │   │       ├── terminal_check_solaris.go
│   │   │       ├── terminal_check_unix.go
│   │   │       ├── terminal_check_windows.go
│   │   │       ├── text_formatter.go
│   │   │       └── writer.go
│   │   ├── skeptycal
│   │   │   ├── util
│   │   │   │   ├── LICENSE
│   │   │   │   ├── gofile
│   │   │   │   │   └── redlogger
│   │   │   │   │       └── redlogger.go
│   │   │   │   └── stringutils
│   │   │   │       ├── LICENSE
│   │   │   │       ├── ansi
│   │   │   │       │   ├── ansi.go
│   │   │   │       │   ├── ansicodes.go
│   │   │   │       │   └── formatstrings.go
│   │   │   │       ├── ascii.go
│   │   │   │       ├── go.doc
│   │   │   │       ├── list.go
│   │   │   │       ├── posix_ascii_man_7.txt
│   │   │   │       ├── set.go
│   │   │   │       ├── stringbenchmarks
│   │   │   │       │   ├── benchmark_results.go
│   │   │   │       │   ├── stringchecks.go
│   │   │   │       │   ├── stringmod.go
│   │   │   │       │   ├── strings.go
│   │   │   │       │   └── whitespace.go
│   │   │   │       ├── strings.go
│   │   │   │       └── valid.go
│   │   │   └── zsh
│   │   │       ├── go.doc
│   │   │       ├── shell.go
│   │   │       ├── sys_calls.go
│   │   │       ├── util.go
│   │   │       └── zsh.go
│   │   └── urfave
│   │       └── cli
│   │           ├── .flake8
│   │           ├── .gitignore
│   │           ├── CODE_OF_CONDUCT.md
│   │           ├── LICENSE
│   │           ├── README.md
│   │           ├── app.go
│   │           ├── appveyor.yml
│   │           ├── category.go
│   │           ├── cli.go
│   │           ├── command.go
│   │           ├── context.go
│   │           ├── docs.go
│   │           ├── errors.go
│   │           ├── fish.go
│   │           ├── flag.go
│   │           ├── flag_bool.go
│   │           ├── flag_bool_t.go
│   │           ├── flag_duration.go
│   │           ├── flag_float64.go
│   │           ├── flag_generic.go
│   │           ├── flag_int.go
│   │           ├── flag_int64.go
│   │           ├── flag_int64_slice.go
│   │           ├── flag_int_slice.go
│   │           ├── flag_string.go
│   │           ├── flag_string_slice.go
│   │           ├── flag_uint.go
│   │           ├── flag_uint64.go
│   │           ├── funcs.go
│   │           ├── help.go
│   │           ├── parse.go
│   │           ├── sort.go
│   │           └── template.go
│   ├── go.opencensus.io
│   │   ├── .gitignore
│   │   ├── .travis.yml
│   │   ├── AUTHORS
│   │   ├── CONTRIBUTING.md
│   │   ├── LICENSE
│   │   ├── Makefile
│   │   ├── README.md
│   │   ├── appveyor.yml
│   │   ├── internal
│   │   │   ├── internal.go
│   │   │   ├── sanitize.go
│   │   │   ├── tagencoding
│   │   │   │   └── tagencoding.go
│   │   │   └── traceinternals.go
│   │   ├── metric
│   │   │   ├── metricdata
│   │   │   │   ├── doc.go
│   │   │   │   ├── exemplar.go
│   │   │   │   ├── label.go
│   │   │   │   ├── metric.go
│   │   │   │   ├── point.go
│   │   │   │   ├── type_string.go
│   │   │   │   └── unit.go
│   │   │   └── metricproducer
│   │   │       ├── manager.go
│   │   │       └── producer.go
│   │   ├── opencensus.go
│   │   ├── plugin
│   │   │   └── ochttp
│   │   │       ├── client.go
│   │   │       ├── client_stats.go
│   │   │       ├── doc.go
│   │   │       ├── propagation
│   │   │       │   └── b3
│   │   │       │       └── b3.go
│   │   │       ├── route.go
│   │   │       ├── server.go
│   │   │       ├── span_annotating_client_trace.go
│   │   │       ├── stats.go
│   │   │       ├── trace.go
│   │   │       └── wrapped_body.go
│   │   ├── resource
│   │   │   └── resource.go
│   │   ├── stats
│   │   │   ├── doc.go
│   │   │   ├── internal
│   │   │   │   └── record.go
│   │   │   ├── measure.go
│   │   │   ├── measure_float64.go
│   │   │   ├── measure_int64.go
│   │   │   ├── record.go
│   │   │   ├── units.go
│   │   │   └── view
│   │   │       ├── aggregation.go
│   │   │       ├── aggregation_data.go
│   │   │       ├── collector.go
│   │   │       ├── doc.go
│   │   │       ├── export.go
│   │   │       ├── view.go
│   │   │       ├── view_to_metric.go
│   │   │       ├── worker.go
│   │   │       └── worker_commands.go
│   │   ├── tag
│   │   │   ├── context.go
│   │   │   ├── doc.go
│   │   │   ├── key.go
│   │   │   ├── map.go
│   │   │   ├── map_codec.go
│   │   │   ├── metadata.go
│   │   │   ├── profile_19.go
│   │   │   ├── profile_not19.go
│   │   │   └── validate.go
│   │   └── trace
│   │       ├── basetypes.go
│   │       ├── config.go
│   │       ├── doc.go
│   │       ├── evictedqueue.go
│   │       ├── export.go
│   │       ├── internal
│   │       │   └── internal.go
│   │       ├── lrumap.go
│   │       ├── propagation
│   │       │   └── propagation.go
│   │       ├── sampling.go
│   │       ├── spanbucket.go
│   │       ├── spanstore.go
│   │       ├── status_codes.go
│   │       ├── trace.go
│   │       ├── trace_go11.go
│   │       ├── trace_nongo11.go
│   │       └── tracestate
│   │           └── tracestate.go
│   ├── golang.org
│   │   └── x
│   │       ├── crypto
│   │       │   ├── AUTHORS
│   │       │   ├── CONTRIBUTORS
│   │       │   ├── LICENSE
│   │       │   ├── PATENTS
│   │       │   ├── bcrypt
│   │       │   │   ├── base64.go
│   │       │   │   └── bcrypt.go
│   │       │   ├── blowfish
│   │       │   │   ├── block.go
│   │       │   │   ├── cipher.go
│   │       │   │   └── const.go
│   │       │   ├── cast5
│   │       │   │   └── cast5.go
│   │       │   └── openpgp
│   │       │       ├── armor
│   │       │       │   ├── armor.go
│   │       │       │   └── encode.go
│   │       │       ├── canonical_text.go
│   │       │       ├── elgamal
│   │       │       │   └── elgamal.go
│   │       │       ├── errors
│   │       │       │   └── errors.go
│   │       │       ├── keys.go
│   │       │       ├── packet
│   │       │       │   ├── compressed.go
│   │       │       │   ├── config.go
│   │       │       │   ├── encrypted_key.go
│   │       │       │   ├── literal.go
│   │       │       │   ├── ocfb.go
│   │       │       │   ├── one_pass_signature.go
│   │       │       │   ├── opaque.go
│   │       │       │   ├── packet.go
│   │       │       │   ├── private_key.go
│   │       │       │   ├── public_key.go
│   │       │       │   ├── public_key_v3.go
│   │       │       │   ├── reader.go
│   │       │       │   ├── signature.go
│   │       │       │   ├── signature_v3.go
│   │       │       │   ├── symmetric_key_encrypted.go
│   │       │       │   ├── symmetrically_encrypted.go
│   │       │       │   ├── userattribute.go
│   │       │       │   └── userid.go
│   │       │       ├── read.go
│   │       │       ├── s2k
│   │       │       │   └── s2k.go
│   │       │       └── write.go
│   │       ├── net
│   │       │   ├── AUTHORS
│   │       │   ├── CONTRIBUTORS
│   │       │   ├── LICENSE
│   │       │   ├── PATENTS
│   │       │   ├── context
│   │       │   │   ├── context.go
│   │       │   │   ├── ctxhttp
│   │       │   │   │   └── ctxhttp.go
│   │       │   │   ├── go17.go
│   │       │   │   ├── go19.go
│   │       │   │   ├── pre_go17.go
│   │       │   │   └── pre_go19.go
│   │       │   ├── http
│   │       │   │   └── httpguts
│   │       │   │       ├── guts.go
│   │       │   │       └── httplex.go
│   │       │   ├── http2
│   │       │   │   ├── .gitignore
│   │       │   │   ├── Dockerfile
│   │       │   │   ├── Makefile
│   │       │   │   ├── ascii.go
│   │       │   │   ├── ciphers.go
│   │       │   │   ├── client_conn_pool.go
│   │       │   │   ├── databuffer.go
│   │       │   │   ├── errors.go
│   │       │   │   ├── flow.go
│   │       │   │   ├── frame.go
│   │       │   │   ├── go111.go
│   │       │   │   ├── go115.go
│   │       │   │   ├── gotrack.go
│   │       │   │   ├── headermap.go
│   │       │   │   ├── hpack
│   │       │   │   │   ├── encode.go
│   │       │   │   │   ├── hpack.go
│   │       │   │   │   ├── huffman.go
│   │       │   │   │   └── tables.go
│   │       │   │   ├── http2.go
│   │       │   │   ├── not_go111.go
│   │       │   │   ├── not_go115.go
│   │       │   │   ├── pipe.go
│   │       │   │   ├── server.go
│   │       │   │   ├── transport.go
│   │       │   │   ├── write.go
│   │       │   │   ├── writesched.go
│   │       │   │   ├── writesched_priority.go
│   │       │   │   └── writesched_random.go
│   │       │   ├── idna
│   │       │   │   ├── go118.go
│   │       │   │   ├── idna10.0.0.go
│   │       │   │   ├── idna9.0.0.go
│   │       │   │   ├── pre_go118.go
│   │       │   │   ├── punycode.go
│   │       │   │   ├── tables10.0.0.go
│   │       │   │   ├── tables11.0.0.go
│   │       │   │   ├── tables12.0.0.go
│   │       │   │   ├── tables13.0.0.go
│   │       │   │   ├── tables9.0.0.go
│   │       │   │   ├── trie.go
│   │       │   │   └── trieval.go
│   │       │   ├── internal
│   │       │   │   └── timeseries
│   │       │   │       └── timeseries.go
│   │       │   └── trace
│   │       │       ├── events.go
│   │       │       ├── histogram.go
│   │       │       └── trace.go
│   │       ├── oauth2
│   │       │   ├── .travis.yml
│   │       │   ├── AUTHORS
│   │       │   ├── CONTRIBUTING.md
│   │       │   ├── CONTRIBUTORS
│   │       │   ├── LICENSE
│   │       │   ├── README.md
│   │       │   ├── authhandler
│   │       │   │   └── authhandler.go
│   │       │   ├── google
│   │       │   │   ├── appengine.go
│   │       │   │   ├── appengine_gen1.go
│   │       │   │   ├── appengine_gen2_flex.go
│   │       │   │   ├── default.go
│   │       │   │   ├── doc.go
│   │       │   │   ├── google.go
│   │       │   │   ├── internal
│   │       │   │   │   └── externalaccount
│   │       │   │   │       ├── aws.go
│   │       │   │   │       ├── basecredentials.go
│   │       │   │   │       ├── clientauth.go
│   │       │   │   │       ├── err.go
│   │       │   │   │       ├── filecredsource.go
│   │       │   │   │       ├── impersonate.go
│   │       │   │   │       ├── sts_exchange.go
│   │       │   │   │       └── urlcredsource.go
│   │       │   │   ├── jwt.go
│   │       │   │   └── sdk.go
│   │       │   ├── internal
│   │       │   │   ├── client_appengine.go
│   │       │   │   ├── doc.go
│   │       │   │   ├── oauth2.go
│   │       │   │   ├── token.go
│   │       │   │   └── transport.go
│   │       │   ├── jws
│   │       │   │   └── jws.go
│   │       │   ├── jwt
│   │       │   │   └── jwt.go
│   │       │   ├── oauth2.go
│   │       │   ├── token.go
│   │       │   └── transport.go
│   │       ├── sync
│   │       │   ├── AUTHORS
│   │       │   ├── CONTRIBUTORS
│   │       │   ├── LICENSE
│   │       │   ├── PATENTS
│   │       │   └── errgroup
│   │       │       └── errgroup.go
│   │       ├── sys
│   │       │   ├── AUTHORS
│   │       │   ├── CONTRIBUTORS
│   │       │   ├── LICENSE
│   │       │   ├── PATENTS
│   │       │   ├── cpu
│   │       │   │   ├── asm_aix_ppc64.s
│   │       │   │   ├── byteorder.go
│   │       │   │   ├── cpu.go
│   │       │   │   ├── cpu_aix.go
│   │       │   │   ├── cpu_arm.go
│   │       │   │   ├── cpu_arm64.go
│   │       │   │   ├── cpu_arm64.s
│   │       │   │   ├── cpu_gc_arm64.go
│   │       │   │   ├── cpu_gc_s390x.go
│   │       │   │   ├── cpu_gc_x86.go
│   │       │   │   ├── cpu_gccgo_arm64.go
│   │       │   │   ├── cpu_gccgo_s390x.go
│   │       │   │   ├── cpu_gccgo_x86.c
│   │       │   │   ├── cpu_gccgo_x86.go
│   │       │   │   ├── cpu_linux.go
│   │       │   │   ├── cpu_linux_arm.go
│   │       │   │   ├── cpu_linux_arm64.go
│   │       │   │   ├── cpu_linux_mips64x.go
│   │       │   │   ├── cpu_linux_noinit.go
│   │       │   │   ├── cpu_linux_ppc64x.go
│   │       │   │   ├── cpu_linux_s390x.go
│   │       │   │   ├── cpu_mips64x.go
│   │       │   │   ├── cpu_mipsx.go
│   │       │   │   ├── cpu_netbsd_arm64.go
│   │       │   │   ├── cpu_other_arm.go
│   │       │   │   ├── cpu_other_arm64.go
│   │       │   │   ├── cpu_other_mips64x.go
│   │       │   │   ├── cpu_ppc64x.go
│   │       │   │   ├── cpu_riscv64.go
│   │       │   │   ├── cpu_s390x.go
│   │       │   │   ├── cpu_s390x.s
│   │       │   │   ├── cpu_wasm.go
│   │       │   │   ├── cpu_x86.go
│   │       │   │   ├── cpu_x86.s
│   │       │   │   ├── cpu_zos.go
│   │       │   │   ├── cpu_zos_s390x.go
│   │       │   │   ├── hwcap_linux.go
│   │       │   │   ├── syscall_aix_gccgo.go
│   │       │   │   └── syscall_aix_ppc64_gc.go
│   │       │   ├── internal
│   │       │   │   └── unsafeheader
│   │       │   │       └── unsafeheader.go
│   │       │   ├── unix
│   │       │   │   ├── .gitignore
│   │       │   │   ├── README.md
│   │       │   │   ├── affinity_linux.go
│   │       │   │   ├── aliases.go
│   │       │   │   ├── asm_aix_ppc64.s
│   │       │   │   ├── asm_bsd_386.s
│   │       │   │   ├── asm_bsd_amd64.s
│   │       │   │   ├── asm_bsd_arm.s
│   │       │   │   ├── asm_bsd_arm64.s
│   │       │   │   ├── asm_linux_386.s
│   │       │   │   ├── asm_linux_amd64.s
│   │       │   │   ├── asm_linux_arm.s
│   │       │   │   ├── asm_linux_arm64.s
│   │       │   │   ├── asm_linux_mips64x.s
│   │       │   │   ├── asm_linux_mipsx.s
│   │       │   │   ├── asm_linux_ppc64x.s
│   │       │   │   ├── asm_linux_riscv64.s
│   │       │   │   ├── asm_linux_s390x.s
│   │       │   │   ├── asm_openbsd_mips64.s
│   │       │   │   ├── asm_solaris_amd64.s
│   │       │   │   ├── asm_zos_s390x.s
│   │       │   │   ├── bluetooth_linux.go
│   │       │   │   ├── cap_freebsd.go
│   │       │   │   ├── constants.go
│   │       │   │   ├── dev_aix_ppc.go
│   │       │   │   ├── dev_aix_ppc64.go
│   │       │   │   ├── dev_darwin.go
│   │       │   │   ├── dev_dragonfly.go
│   │       │   │   ├── dev_freebsd.go
│   │       │   │   ├── dev_linux.go
│   │       │   │   ├── dev_netbsd.go
│   │       │   │   ├── dev_openbsd.go
│   │       │   │   ├── dev_zos.go
│   │       │   │   ├── dirent.go
│   │       │   │   ├── endian_big.go
│   │       │   │   ├── endian_little.go
│   │       │   │   ├── env_unix.go
│   │       │   │   ├── epoll_zos.go
│   │       │   │   ├── errors_freebsd_386.go
│   │       │   │   ├── errors_freebsd_amd64.go
│   │       │   │   ├── errors_freebsd_arm.go
│   │       │   │   ├── errors_freebsd_arm64.go
│   │       │   │   ├── fcntl.go
│   │       │   │   ├── fcntl_darwin.go
│   │       │   │   ├── fcntl_linux_32bit.go
│   │       │   │   ├── fdset.go
│   │       │   │   ├── fstatfs_zos.go
│   │       │   │   ├── gccgo.go
│   │       │   │   ├── gccgo_c.c
│   │       │   │   ├── gccgo_linux_amd64.go
│   │       │   │   ├── ifreq_linux.go
│   │       │   │   ├── ioctl.go
│   │       │   │   ├── ioctl_linux.go
│   │       │   │   ├── ioctl_zos.go
│   │       │   │   ├── mkall.sh
│   │       │   │   ├── mkerrors.sh
│   │       │   │   ├── pagesize_unix.go
│   │       │   │   ├── pledge_openbsd.go
│   │       │   │   ├── ptrace_darwin.go
│   │       │   │   ├── ptrace_ios.go
│   │       │   │   ├── race.go
│   │       │   │   ├── race0.go
│   │       │   │   ├── readdirent_getdents.go
│   │       │   │   ├── readdirent_getdirentries.go
│   │       │   │   ├── sockcmsg_dragonfly.go
│   │       │   │   ├── sockcmsg_linux.go
│   │       │   │   ├── sockcmsg_unix.go
│   │       │   │   ├── sockcmsg_unix_other.go
│   │       │   │   ├── str.go
│   │       │   │   ├── syscall.go
│   │       │   │   ├── syscall_aix.go
│   │       │   │   ├── syscall_aix_ppc.go
│   │       │   │   ├── syscall_aix_ppc64.go
│   │       │   │   ├── syscall_bsd.go
│   │       │   │   ├── syscall_darwin.1_12.go
│   │       │   │   ├── syscall_darwin.1_13.go
│   │       │   │   ├── syscall_darwin.go
│   │       │   │   ├── syscall_darwin_amd64.go
│   │       │   │   ├── syscall_darwin_arm64.go
│   │       │   │   ├── syscall_darwin_libSystem.go
│   │       │   │   ├── syscall_dragonfly.go
│   │       │   │   ├── syscall_dragonfly_amd64.go
│   │       │   │   ├── syscall_freebsd.go
│   │       │   │   ├── syscall_freebsd_386.go
│   │       │   │   ├── syscall_freebsd_amd64.go
│   │       │   │   ├── syscall_freebsd_arm.go
│   │       │   │   ├── syscall_freebsd_arm64.go
│   │       │   │   ├── syscall_illumos.go
│   │       │   │   ├── syscall_linux.go
│   │       │   │   ├── syscall_linux_386.go
│   │       │   │   ├── syscall_linux_amd64.go
│   │       │   │   ├── syscall_linux_amd64_gc.go
│   │       │   │   ├── syscall_linux_arm.go
│   │       │   │   ├── syscall_linux_arm64.go
│   │       │   │   ├── syscall_linux_gc.go
│   │       │   │   ├── syscall_linux_gc_386.go
│   │       │   │   ├── syscall_linux_gc_arm.go
│   │       │   │   ├── syscall_linux_gccgo_386.go
│   │       │   │   ├── syscall_linux_gccgo_arm.go
│   │       │   │   ├── syscall_linux_mips64x.go
│   │       │   │   ├── syscall_linux_mipsx.go
│   │       │   │   ├── syscall_linux_ppc.go
│   │       │   │   ├── syscall_linux_ppc64x.go
│   │       │   │   ├── syscall_linux_riscv64.go
│   │       │   │   ├── syscall_linux_s390x.go
│   │       │   │   ├── syscall_linux_sparc64.go
│   │       │   │   ├── syscall_netbsd.go
│   │       │   │   ├── syscall_netbsd_386.go
│   │       │   │   ├── syscall_netbsd_amd64.go
│   │       │   │   ├── syscall_netbsd_arm.go
│   │       │   │   ├── syscall_netbsd_arm64.go
│   │       │   │   ├── syscall_openbsd.go
│   │       │   │   ├── syscall_openbsd_386.go
│   │       │   │   ├── syscall_openbsd_amd64.go
│   │       │   │   ├── syscall_openbsd_arm.go
│   │       │   │   ├── syscall_openbsd_arm64.go
│   │       │   │   ├── syscall_openbsd_mips64.go
│   │       │   │   ├── syscall_solaris.go
│   │       │   │   ├── syscall_solaris_amd64.go
│   │       │   │   ├── syscall_unix.go
│   │       │   │   ├── syscall_unix_gc.go
│   │       │   │   ├── syscall_unix_gc_ppc64x.go
│   │       │   │   ├── syscall_zos_s390x.go
│   │       │   │   ├── sysvshm_linux.go
│   │       │   │   ├── sysvshm_unix.go
│   │       │   │   ├── sysvshm_unix_other.go
│   │       │   │   ├── timestruct.go
│   │       │   │   ├── unveil_openbsd.go
│   │       │   │   ├── xattr_bsd.go
│   │       │   │   ├── zerrors_aix_ppc.go
│   │       │   │   ├── zerrors_aix_ppc64.go
│   │       │   │   ├── zerrors_darwin_amd64.go
│   │       │   │   ├── zerrors_darwin_arm64.go
│   │       │   │   ├── zerrors_dragonfly_amd64.go
│   │       │   │   ├── zerrors_freebsd_386.go
│   │       │   │   ├── zerrors_freebsd_amd64.go
│   │       │   │   ├── zerrors_freebsd_arm.go
│   │       │   │   ├── zerrors_freebsd_arm64.go
│   │       │   │   ├── zerrors_linux.go
│   │       │   │   ├── zerrors_linux_386.go
│   │       │   │   ├── zerrors_linux_amd64.go
│   │       │   │   ├── zerrors_linux_arm.go
│   │       │   │   ├── zerrors_linux_arm64.go
│   │       │   │   ├── zerrors_linux_mips.go
│   │       │   │   ├── zerrors_linux_mips64.go
│   │       │   │   ├── zerrors_linux_mips64le.go
│   │       │   │   ├── zerrors_linux_mipsle.go
│   │       │   │   ├── zerrors_linux_ppc.go
│   │       │   │   ├── zerrors_linux_ppc64.go
│   │       │   │   ├── zerrors_linux_ppc64le.go
│   │       │   │   ├── zerrors_linux_riscv64.go
│   │       │   │   ├── zerrors_linux_s390x.go
│   │       │   │   ├── zerrors_linux_sparc64.go
│   │       │   │   ├── zerrors_netbsd_386.go
│   │       │   │   ├── zerrors_netbsd_amd64.go
│   │       │   │   ├── zerrors_netbsd_arm.go
│   │       │   │   ├── zerrors_netbsd_arm64.go
│   │       │   │   ├── zerrors_openbsd_386.go
│   │       │   │   ├── zerrors_openbsd_amd64.go
│   │       │   │   ├── zerrors_openbsd_arm.go
│   │       │   │   ├── zerrors_openbsd_arm64.go
│   │       │   │   ├── zerrors_openbsd_mips64.go
│   │       │   │   ├── zerrors_solaris_amd64.go
│   │       │   │   ├── zerrors_zos_s390x.go
│   │       │   │   ├── zptrace_armnn_linux.go
│   │       │   │   ├── zptrace_linux_arm64.go
│   │       │   │   ├── zptrace_mipsnn_linux.go
│   │       │   │   ├── zptrace_mipsnnle_linux.go
│   │       │   │   ├── zptrace_x86_linux.go
│   │       │   │   ├── zsyscall_aix_ppc.go
│   │       │   │   ├── zsyscall_aix_ppc64.go
│   │       │   │   ├── zsyscall_aix_ppc64_gc.go
│   │       │   │   ├── zsyscall_aix_ppc64_gccgo.go
│   │       │   │   ├── zsyscall_darwin_amd64.1_13.go
│   │       │   │   ├── zsyscall_darwin_amd64.1_13.s
│   │       │   │   ├── zsyscall_darwin_amd64.go
│   │       │   │   ├── zsyscall_darwin_amd64.s
│   │       │   │   ├── zsyscall_darwin_arm64.1_13.go
│   │       │   │   ├── zsyscall_darwin_arm64.1_13.s
│   │       │   │   ├── zsyscall_darwin_arm64.go
│   │       │   │   ├── zsyscall_darwin_arm64.s
│   │       │   │   ├── zsyscall_dragonfly_amd64.go
│   │       │   │   ├── zsyscall_freebsd_386.go
│   │       │   │   ├── zsyscall_freebsd_amd64.go
│   │       │   │   ├── zsyscall_freebsd_arm.go
│   │       │   │   ├── zsyscall_freebsd_arm64.go
│   │       │   │   ├── zsyscall_illumos_amd64.go
│   │       │   │   ├── zsyscall_linux.go
│   │       │   │   ├── zsyscall_linux_386.go
│   │       │   │   ├── zsyscall_linux_amd64.go
│   │       │   │   ├── zsyscall_linux_arm.go
│   │       │   │   ├── zsyscall_linux_arm64.go
│   │       │   │   ├── zsyscall_linux_mips.go
│   │       │   │   ├── zsyscall_linux_mips64.go
│   │       │   │   ├── zsyscall_linux_mips64le.go
│   │       │   │   ├── zsyscall_linux_mipsle.go
│   │       │   │   ├── zsyscall_linux_ppc.go
│   │       │   │   ├── zsyscall_linux_ppc64.go
│   │       │   │   ├── zsyscall_linux_ppc64le.go
│   │       │   │   ├── zsyscall_linux_riscv64.go
│   │       │   │   ├── zsyscall_linux_s390x.go
│   │       │   │   ├── zsyscall_linux_sparc64.go
│   │       │   │   ├── zsyscall_netbsd_386.go
│   │       │   │   ├── zsyscall_netbsd_amd64.go
│   │       │   │   ├── zsyscall_netbsd_arm.go
│   │       │   │   ├── zsyscall_netbsd_arm64.go
│   │       │   │   ├── zsyscall_openbsd_386.go
│   │       │   │   ├── zsyscall_openbsd_amd64.go
│   │       │   │   ├── zsyscall_openbsd_arm.go
│   │       │   │   ├── zsyscall_openbsd_arm64.go
│   │       │   │   ├── zsyscall_openbsd_mips64.go
│   │       │   │   ├── zsyscall_solaris_amd64.go
│   │       │   │   ├── zsyscall_zos_s390x.go
│   │       │   │   ├── zsysctl_openbsd_386.go
│   │       │   │   ├── zsysctl_openbsd_amd64.go
│   │       │   │   ├── zsysctl_openbsd_arm.go
│   │       │   │   ├── zsysctl_openbsd_arm64.go
│   │       │   │   ├── zsysctl_openbsd_mips64.go
│   │       │   │   ├── zsysnum_darwin_amd64.go
│   │       │   │   ├── zsysnum_darwin_arm64.go
│   │       │   │   ├── zsysnum_dragonfly_amd64.go
│   │       │   │   ├── zsysnum_freebsd_386.go
│   │       │   │   ├── zsysnum_freebsd_amd64.go
│   │       │   │   ├── zsysnum_freebsd_arm.go
│   │       │   │   ├── zsysnum_freebsd_arm64.go
│   │       │   │   ├── zsysnum_linux_386.go
│   │       │   │   ├── zsysnum_linux_amd64.go
│   │       │   │   ├── zsysnum_linux_arm.go
│   │       │   │   ├── zsysnum_linux_arm64.go
│   │       │   │   ├── zsysnum_linux_mips.go
│   │       │   │   ├── zsysnum_linux_mips64.go
│   │       │   │   ├── zsysnum_linux_mips64le.go
│   │       │   │   ├── zsysnum_linux_mipsle.go
│   │       │   │   ├── zsysnum_linux_ppc.go
│   │       │   │   ├── zsysnum_linux_ppc64.go
│   │       │   │   ├── zsysnum_linux_ppc64le.go
│   │       │   │   ├── zsysnum_linux_riscv64.go
│   │       │   │   ├── zsysnum_linux_s390x.go
│   │       │   │   ├── zsysnum_linux_sparc64.go
│   │       │   │   ├── zsysnum_netbsd_386.go
│   │       │   │   ├── zsysnum_netbsd_amd64.go
│   │       │   │   ├── zsysnum_netbsd_arm.go
│   │       │   │   ├── zsysnum_netbsd_arm64.go
│   │       │   │   ├── zsysnum_openbsd_386.go
│   │       │   │   ├── zsysnum_openbsd_amd64.go
│   │       │   │   ├── zsysnum_openbsd_arm.go
│   │       │   │   ├── zsysnum_openbsd_arm64.go
│   │       │   │   ├── zsysnum_openbsd_mips64.go
│   │       │   │   ├── zsysnum_zos_s390x.go
│   │       │   │   ├── ztypes_aix_ppc.go
│   │       │   │   ├── ztypes_aix_ppc64.go
│   │       │   │   ├── ztypes_darwin_amd64.go
│   │       │   │   ├── ztypes_darwin_arm64.go
│   │       │   │   ├── ztypes_dragonfly_amd64.go
│   │       │   │   ├── ztypes_freebsd_386.go
│   │       │   │   ├── ztypes_freebsd_amd64.go
│   │       │   │   ├── ztypes_freebsd_arm.go
│   │       │   │   ├── ztypes_freebsd_arm64.go
│   │       │   │   ├── ztypes_illumos_amd64.go
│   │       │   │   ├── ztypes_linux.go
│   │       │   │   ├── ztypes_linux_386.go
│   │       │   │   ├── ztypes_linux_amd64.go
│   │       │   │   ├── ztypes_linux_arm.go
│   │       │   │   ├── ztypes_linux_arm64.go
│   │       │   │   ├── ztypes_linux_mips.go
│   │       │   │   ├── ztypes_linux_mips64.go
│   │       │   │   ├── ztypes_linux_mips64le.go
│   │       │   │   ├── ztypes_linux_mipsle.go
│   │       │   │   ├── ztypes_linux_ppc.go
│   │       │   │   ├── ztypes_linux_ppc64.go
│   │       │   │   ├── ztypes_linux_ppc64le.go
│   │       │   │   ├── ztypes_linux_riscv64.go
│   │       │   │   ├── ztypes_linux_s390x.go
│   │       │   │   ├── ztypes_linux_sparc64.go
│   │       │   │   ├── ztypes_netbsd_386.go
│   │       │   │   ├── ztypes_netbsd_amd64.go
│   │       │   │   ├── ztypes_netbsd_arm.go
│   │       │   │   ├── ztypes_netbsd_arm64.go
│   │       │   │   ├── ztypes_openbsd_386.go
│   │       │   │   ├── ztypes_openbsd_amd64.go
│   │       │   │   ├── ztypes_openbsd_arm.go
│   │       │   │   ├── ztypes_openbsd_arm64.go
│   │       │   │   ├── ztypes_openbsd_mips64.go
│   │       │   │   ├── ztypes_solaris_amd64.go
│   │       │   │   └── ztypes_zos_s390x.go
│   │       │   └── windows
│   │       │       ├── aliases.go
│   │       │       ├── dll_windows.go
│   │       │       ├── empty.s
│   │       │       ├── env_windows.go
│   │       │       ├── eventlog.go
│   │       │       ├── exec_windows.go
│   │       │       ├── memory_windows.go
│   │       │       ├── mkerrors.bash
│   │       │       ├── mkknownfolderids.bash
│   │       │       ├── mksyscall.go
│   │       │       ├── race.go
│   │       │       ├── race0.go
│   │       │       ├── security_windows.go
│   │       │       ├── service.go
│   │       │       ├── setupapi_windows.go
│   │       │       ├── str.go
│   │       │       ├── syscall.go
│   │       │       ├── syscall_windows.go
│   │       │       ├── types_windows.go
│   │       │       ├── types_windows_386.go
│   │       │       ├── types_windows_amd64.go
│   │       │       ├── types_windows_arm.go
│   │       │       ├── types_windows_arm64.go
│   │       │       ├── zerrors_windows.go
│   │       │       ├── zknownfolderids_windows.go
│   │       │       └── zsyscall_windows.go
│   │       └── text
│   │           ├── AUTHORS
│   │           ├── CONTRIBUTORS
│   │           ├── LICENSE
│   │           ├── PATENTS
│   │           ├── secure
│   │           │   └── bidirule
│   │           │       ├── bidirule.go
│   │           │       ├── bidirule10.0.0.go
│   │           │       └── bidirule9.0.0.go
│   │           ├── transform
│   │           │   └── transform.go
│   │           └── unicode
│   │               ├── bidi
│   │               │   ├── bidi.go
│   │               │   ├── bracket.go
│   │               │   ├── core.go
│   │               │   ├── prop.go
│   │               │   ├── tables10.0.0.go
│   │               │   ├── tables11.0.0.go
│   │               │   ├── tables12.0.0.go
│   │               │   ├── tables13.0.0.go
│   │               │   ├── tables9.0.0.go
│   │               │   └── trieval.go
│   │               └── norm
│   │                   ├── composition.go
│   │                   ├── forminfo.go
│   │                   ├── input.go
│   │                   ├── iter.go
│   │                   ├── normalize.go
│   │                   ├── readwriter.go
│   │                   ├── tables10.0.0.go
│   │                   ├── tables11.0.0.go
│   │                   ├── tables12.0.0.go
│   │                   ├── tables13.0.0.go
│   │                   ├── tables9.0.0.go
│   │                   ├── transform.go
│   │                   └── trie.go
│   ├── google.golang.org
│   │   ├── api
│   │   │   ├── AUTHORS
│   │   │   ├── CONTRIBUTORS
│   │   │   ├── LICENSE
│   │   │   ├── gmail
│   │   │   │   └── v1
│   │   │   │       ├── gmail-api.json
│   │   │   │       └── gmail-gen.go
│   │   │   ├── googleapi
│   │   │   │   ├── googleapi.go
│   │   │   │   ├── transport
│   │   │   │   │   └── apikey.go
│   │   │   │   └── types.go
│   │   │   ├── internal
│   │   │   │   ├── conn_pool.go
│   │   │   │   ├── creds.go
│   │   │   │   ├── gensupport
│   │   │   │   │   ├── buffer.go
│   │   │   │   │   ├── doc.go
│   │   │   │   │   ├── json.go
│   │   │   │   │   ├── jsonfloat.go
│   │   │   │   │   ├── media.go
│   │   │   │   │   ├── params.go
│   │   │   │   │   ├── resumable.go
│   │   │   │   │   ├── retryable_linux.go
│   │   │   │   │   ├── send.go
│   │   │   │   │   └── version.go
│   │   │   │   ├── service-account.json
│   │   │   │   ├── settings.go
│   │   │   │   └── third_party
│   │   │   │       └── uritemplates
│   │   │   │           ├── LICENSE
│   │   │   │           ├── METADATA
│   │   │   │           ├── uritemplates.go
│   │   │   │           └── utils.go
│   │   │   ├── option
│   │   │   │   ├── credentials_go19.go
│   │   │   │   ├── credentials_notgo19.go
│   │   │   │   ├── internaloption
│   │   │   │   │   └── internaloption.go
│   │   │   │   └── option.go
│   │   │   └── transport
│   │   │       ├── cert
│   │   │       │   └── default_cert.go
│   │   │       └── http
│   │   │           ├── default_transport_go113.go
│   │   │           ├── default_transport_not_go113.go
│   │   │           ├── dial.go
│   │   │           ├── dial_appengine.go
│   │   │           └── internal
│   │   │               └── propagation
│   │   │                   └── http.go
│   │   ├── appengine
│   │   │   ├── .travis.yml
│   │   │   ├── CONTRIBUTING.md
│   │   │   ├── LICENSE
│   │   │   ├── README.md
│   │   │   ├── appengine.go
│   │   │   ├── appengine_vm.go
│   │   │   ├── errors.go
│   │   │   ├── identity.go
│   │   │   ├── internal
│   │   │   │   ├── api.go
│   │   │   │   ├── api_classic.go
│   │   │   │   ├── api_common.go
│   │   │   │   ├── app_id.go
│   │   │   │   ├── app_identity
│   │   │   │   │   ├── app_identity_service.pb.go
│   │   │   │   │   └── app_identity_service.proto
│   │   │   │   ├── base
│   │   │   │   │   ├── api_base.pb.go
│   │   │   │   │   └── api_base.proto
│   │   │   │   ├── datastore
│   │   │   │   │   ├── datastore_v3.pb.go
│   │   │   │   │   └── datastore_v3.proto
│   │   │   │   ├── identity.go
│   │   │   │   ├── identity_classic.go
│   │   │   │   ├── identity_flex.go
│   │   │   │   ├── identity_vm.go
│   │   │   │   ├── internal.go
│   │   │   │   ├── log
│   │   │   │   │   ├── log_service.pb.go
│   │   │   │   │   └── log_service.proto
│   │   │   │   ├── main.go
│   │   │   │   ├── main_common.go
│   │   │   │   ├── main_vm.go
│   │   │   │   ├── metadata.go
│   │   │   │   ├── modules
│   │   │   │   │   ├── modules_service.pb.go
│   │   │   │   │   └── modules_service.proto
│   │   │   │   ├── net.go
│   │   │   │   ├── regen.sh
│   │   │   │   ├── remote_api
│   │   │   │   │   ├── remote_api.pb.go
│   │   │   │   │   └── remote_api.proto
│   │   │   │   ├── transaction.go
│   │   │   │   └── urlfetch
│   │   │   │       ├── urlfetch_service.pb.go
│   │   │   │       └── urlfetch_service.proto
│   │   │   ├── namespace.go
│   │   │   ├── timeout.go
│   │   │   ├── travis_install.sh
│   │   │   ├── travis_test.sh
│   │   │   └── urlfetch
│   │   │       └── urlfetch.go
│   │   ├── genproto
│   │   │   ├── LICENSE
│   │   │   └── googleapis
│   │   │       └── rpc
│   │   │           └── status
│   │   │               └── status.pb.go
│   │   ├── grpc
│   │   │   ├── .travis.yml
│   │   │   ├── AUTHORS
│   │   │   ├── CODE-OF-CONDUCT.md
│   │   │   ├── CONTRIBUTING.md
│   │   │   ├── GOVERNANCE.md
│   │   │   ├── LICENSE
│   │   │   ├── MAINTAINERS.md
│   │   │   ├── Makefile
│   │   │   ├── README.md
│   │   │   ├── attributes
│   │   │   │   └── attributes.go
│   │   │   ├── backoff
│   │   │   │   └── backoff.go
│   │   │   ├── backoff.go
│   │   │   ├── balancer
│   │   │   │   ├── balancer.go
│   │   │   │   ├── base
│   │   │   │   │   ├── balancer.go
│   │   │   │   │   └── base.go
│   │   │   │   ├── grpclb
│   │   │   │   │   └── state
│   │   │   │   │       └── state.go
│   │   │   │   └── roundrobin
│   │   │   │       └── roundrobin.go
│   │   │   ├── balancer_conn_wrappers.go
│   │   │   ├── binarylog
│   │   │   │   └── grpc_binarylog_v1
│   │   │   │       └── binarylog.pb.go
│   │   │   ├── call.go
│   │   │   ├── clientconn.go
│   │   │   ├── codec.go
│   │   │   ├── codegen.sh
│   │   │   ├── codes
│   │   │   │   ├── code_string.go
│   │   │   │   └── codes.go
│   │   │   ├── connectivity
│   │   │   │   └── connectivity.go
│   │   │   ├── credentials
│   │   │   │   ├── credentials.go
│   │   │   │   ├── go12.go
│   │   │   │   ├── internal
│   │   │   │   │   ├── syscallconn.go
│   │   │   │   │   └── syscallconn_appengine.go
│   │   │   │   └── tls.go
│   │   │   ├── dialoptions.go
│   │   │   ├── doc.go
│   │   │   ├── encoding
│   │   │   │   ├── encoding.go
│   │   │   │   └── proto
│   │   │   │       └── proto.go
│   │   │   ├── grpclog
│   │   │   │   ├── component.go
│   │   │   │   ├── grpclog.go
│   │   │   │   ├── logger.go
│   │   │   │   └── loggerv2.go
│   │   │   ├── install_gae.sh
│   │   │   ├── interceptor.go
│   │   │   ├── internal
│   │   │   │   ├── backoff
│   │   │   │   │   └── backoff.go
│   │   │   │   ├── balancerload
│   │   │   │   │   └── load.go
│   │   │   │   ├── binarylog
│   │   │   │   │   ├── binarylog.go
│   │   │   │   │   ├── binarylog_testutil.go
│   │   │   │   │   ├── env_config.go
│   │   │   │   │   ├── method_logger.go
│   │   │   │   │   └── sink.go
│   │   │   │   ├── buffer
│   │   │   │   │   └── unbounded.go
│   │   │   │   ├── channelz
│   │   │   │   │   ├── funcs.go
│   │   │   │   │   ├── logging.go
│   │   │   │   │   ├── types.go
│   │   │   │   │   ├── types_linux.go
│   │   │   │   │   ├── types_nonlinux.go
│   │   │   │   │   ├── util_linux.go
│   │   │   │   │   └── util_nonlinux.go
│   │   │   │   ├── credentials
│   │   │   │   │   ├── go110.go
│   │   │   │   │   └── gobefore110.go
│   │   │   │   ├── envconfig
│   │   │   │   │   └── envconfig.go
│   │   │   │   ├── grpclog
│   │   │   │   │   ├── grpclog.go
│   │   │   │   │   └── prefixLogger.go
│   │   │   │   ├── grpcrand
│   │   │   │   │   └── grpcrand.go
│   │   │   │   ├── grpcsync
│   │   │   │   │   └── event.go
│   │   │   │   ├── grpcutil
│   │   │   │   │   ├── encode_duration.go
│   │   │   │   │   ├── metadata.go
│   │   │   │   │   ├── method.go
│   │   │   │   │   └── target.go
│   │   │   │   ├── internal.go
│   │   │   │   ├── resolver
│   │   │   │   │   ├── dns
│   │   │   │   │   │   ├── dns_resolver.go
│   │   │   │   │   │   └── go113.go
│   │   │   │   │   └── passthrough
│   │   │   │   │       └── passthrough.go
│   │   │   │   ├── serviceconfig
│   │   │   │   │   └── serviceconfig.go
│   │   │   │   ├── status
│   │   │   │   │   └── status.go
│   │   │   │   ├── syscall
│   │   │   │   │   ├── syscall_linux.go
│   │   │   │   │   └── syscall_nonlinux.go
│   │   │   │   └── transport
│   │   │   │       ├── bdp_estimator.go
│   │   │   │       ├── controlbuf.go
│   │   │   │       ├── defaults.go
│   │   │   │       ├── flowcontrol.go
│   │   │   │       ├── handler_server.go
│   │   │   │       ├── http2_client.go
│   │   │   │       ├── http2_server.go
│   │   │   │       ├── http_util.go
│   │   │   │       └── transport.go
│   │   │   ├── keepalive
│   │   │   │   └── keepalive.go
│   │   │   ├── metadata
│   │   │   │   └── metadata.go
│   │   │   ├── peer
│   │   │   │   └── peer.go
│   │   │   ├── picker_wrapper.go
│   │   │   ├── pickfirst.go
│   │   │   ├── preloader.go
│   │   │   ├── proxy.go
│   │   │   ├── regenerate.sh
│   │   │   ├── resolver
│   │   │   │   └── resolver.go
│   │   │   ├── resolver_conn_wrapper.go
│   │   │   ├── rpc_util.go
│   │   │   ├── server.go
│   │   │   ├── service_config.go
│   │   │   ├── serviceconfig
│   │   │   │   └── serviceconfig.go
│   │   │   ├── stats
│   │   │   │   ├── handlers.go
│   │   │   │   └── stats.go
│   │   │   ├── status
│   │   │   │   └── status.go
│   │   │   ├── stream.go
│   │   │   ├── tap
│   │   │   │   └── tap.go
│   │   │   ├── trace.go
│   │   │   ├── version.go
│   │   │   └── vet.sh
│   │   └── protobuf
│   │       ├── AUTHORS
│   │       ├── CONTRIBUTORS
│   │       ├── LICENSE
│   │       ├── PATENTS
│   │       ├── encoding
│   │       │   ├── prototext
│   │       │   │   ├── decode.go
│   │       │   │   ├── doc.go
│   │       │   │   └── encode.go
│   │       │   └── protowire
│   │       │       └── wire.go
│   │       ├── internal
│   │       │   ├── descfmt
│   │       │   │   └── stringer.go
│   │       │   ├── descopts
│   │       │   │   └── options.go
│   │       │   ├── detrand
│   │       │   │   └── rand.go
│   │       │   ├── encoding
│   │       │   │   ├── defval
│   │       │   │   │   └── default.go
│   │       │   │   ├── messageset
│   │       │   │   │   └── messageset.go
│   │       │   │   ├── tag
│   │       │   │   │   └── tag.go
│   │       │   │   └── text
│   │       │   │       ├── decode.go
│   │       │   │       ├── decode_number.go
│   │       │   │       ├── decode_string.go
│   │       │   │       ├── decode_token.go
│   │       │   │       ├── doc.go
│   │       │   │       └── encode.go
│   │       │   ├── errors
│   │       │   │   ├── errors.go
│   │       │   │   ├── is_go112.go
│   │       │   │   └── is_go113.go
│   │       │   ├── fieldsort
│   │       │   │   └── fieldsort.go
│   │       │   ├── filedesc
│   │       │   │   ├── build.go
│   │       │   │   ├── desc.go
│   │       │   │   ├── desc_init.go
│   │       │   │   ├── desc_lazy.go
│   │       │   │   ├── desc_list.go
│   │       │   │   ├── desc_list_gen.go
│   │       │   │   └── placeholder.go
│   │       │   ├── filetype
│   │       │   │   └── build.go
│   │       │   ├── flags
│   │       │   │   ├── flags.go
│   │       │   │   ├── proto_legacy_disable.go
│   │       │   │   └── proto_legacy_enable.go
│   │       │   ├── genid
│   │       │   │   ├── any_gen.go
│   │       │   │   ├── api_gen.go
│   │       │   │   ├── descriptor_gen.go
│   │       │   │   ├── doc.go
│   │       │   │   ├── duration_gen.go
│   │       │   │   ├── empty_gen.go
│   │       │   │   ├── field_mask_gen.go
│   │       │   │   ├── goname.go
│   │       │   │   ├── map_entry.go
│   │       │   │   ├── source_context_gen.go
│   │       │   │   ├── struct_gen.go
│   │       │   │   ├── timestamp_gen.go
│   │       │   │   ├── type_gen.go
│   │       │   │   ├── wrappers.go
│   │       │   │   └── wrappers_gen.go
│   │       │   ├── impl
│   │       │   │   ├── api_export.go
│   │       │   │   ├── checkinit.go
│   │       │   │   ├── codec_extension.go
│   │       │   │   ├── codec_field.go
│   │       │   │   ├── codec_gen.go
│   │       │   │   ├── codec_map.go
│   │       │   │   ├── codec_map_go111.go
│   │       │   │   ├── codec_map_go112.go
│   │       │   │   ├── codec_message.go
│   │       │   │   ├── codec_messageset.go
│   │       │   │   ├── codec_reflect.go
│   │       │   │   ├── codec_tables.go
│   │       │   │   ├── codec_unsafe.go
│   │       │   │   ├── convert.go
│   │       │   │   ├── convert_list.go
│   │       │   │   ├── convert_map.go
│   │       │   │   ├── decode.go
│   │       │   │   ├── encode.go
│   │       │   │   ├── enum.go
│   │       │   │   ├── extension.go
│   │       │   │   ├── legacy_enum.go
│   │       │   │   ├── legacy_export.go
│   │       │   │   ├── legacy_extension.go
│   │       │   │   ├── legacy_file.go
│   │       │   │   ├── legacy_message.go
│   │       │   │   ├── merge.go
│   │       │   │   ├── merge_gen.go
│   │       │   │   ├── message.go
│   │       │   │   ├── message_reflect.go
│   │       │   │   ├── message_reflect_field.go
│   │       │   │   ├── message_reflect_gen.go
│   │       │   │   ├── pointer_reflect.go
│   │       │   │   ├── pointer_unsafe.go
│   │       │   │   ├── validate.go
│   │       │   │   └── weak.go
│   │       │   ├── mapsort
│   │       │   │   └── mapsort.go
│   │       │   ├── pragma
│   │       │   │   └── pragma.go
│   │       │   ├── set
│   │       │   │   └── ints.go
│   │       │   ├── strs
│   │       │   │   ├── strings.go
│   │       │   │   ├── strings_pure.go
│   │       │   │   └── strings_unsafe.go
│   │       │   └── version
│   │       │       └── version.go
│   │       ├── proto
│   │       │   ├── checkinit.go
│   │       │   ├── decode.go
│   │       │   ├── decode_gen.go
│   │       │   ├── doc.go
│   │       │   ├── encode.go
│   │       │   ├── encode_gen.go
│   │       │   ├── equal.go
│   │       │   ├── extension.go
│   │       │   ├── merge.go
│   │       │   ├── messageset.go
│   │       │   ├── proto.go
│   │       │   ├── proto_methods.go
│   │       │   ├── proto_reflect.go
│   │       │   ├── reset.go
│   │       │   ├── size.go
│   │       │   ├── size_gen.go
│   │       │   └── wrappers.go
│   │       ├── reflect
│   │       │   ├── protoreflect
│   │       │   │   ├── methods.go
│   │       │   │   ├── proto.go
│   │       │   │   ├── source.go
│   │       │   │   ├── type.go
│   │       │   │   ├── value.go
│   │       │   │   ├── value_pure.go
│   │       │   │   ├── value_union.go
│   │       │   │   └── value_unsafe.go
│   │       │   └── protoregistry
│   │       │       └── registry.go
│   │       ├── runtime
│   │       │   ├── protoiface
│   │       │   │   ├── legacy.go
│   │       │   │   └── methods.go
│   │       │   └── protoimpl
│   │       │       ├── impl.go
│   │       │       └── version.go
│   │       └── types
│   │           └── known
│   │               ├── anypb
│   │               │   └── any.pb.go
│   │               ├── durationpb
│   │               │   └── duration.pb.go
│   │               └── timestamppb
│   │                   └── timestamp.pb.go
│   └── modules.txt
├── vibe
│   └── idea.md
└── webtools
    ├── LICENSE
    ├── getpage
    │   └── googlesearch.css
    ├── http
    │   └── downloadurl.go
    └── youtube
        └── gotube
            ├── gotube
            └── main.go

424 directories, 1888 files
```

[get_tree]: (http://mama.indstate.edu/users/ice/tree/)
