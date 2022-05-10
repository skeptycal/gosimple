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
│   ├── size
│   │   ├── size.go
│   │   └── size_test.go
│   └── structures
│       ├── array
│       │   └── array.go
│       ├── arraylist
│       │   └── arraylist.go
│       ├── bst
│       │   └── bst.go
│       ├── dict
│       │   └── dict.go
│       ├── doublylinkedlist
│       │   └── doublylinkedlist.go
│       ├── linkedlist
│       │   └── linkedlist.go
│       ├── queue
│       │   └── queue.go
│       ├── stack
│       │   ├── example
│       │   │   └── main.go
│       │   └── stack.go
│       ├── structures.go
│       └── tree
│           └── tree.go
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
├── profile1652215597629343000.out
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
│   │   │   ├── cli
│   │   │   │   └── cli.go
│   │   │   ├── file
│   │   │   │   ├── bytes.go
│   │   │   │   ├── const.go
│   │   │   │   ├── constants.go
│   │   │   │   ├── debug.go
│   │   │   │   ├── file.go
│   │   │   │   ├── filebasic.go
│   │   │   │   ├── fileops.go
│   │   │   │   ├── flags.go
│   │   │   │   ├── imports.go
│   │   │   │   ├── lines.go
│   │   │   │   └── memfile.go
│   │   │   ├── gilist.bak.txt
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

193 directories, 660 files
```

[get_tree]: (http://mama.indstate.edu/users/ice/tree/)
