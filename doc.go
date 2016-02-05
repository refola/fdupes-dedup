// Copyright 2015 Mark Haferkamp. This code is licensed under the
// license found in the LICENSE file.

package doc

// fdupes-dedupe is a program that deduplicates files in copy-on-write
// (CoW) filesystems by deleting duplicate files and recreating them
// as CoW copies. This describes how it works in general.

/*
Outline:

* Get dedupe path from user (-path=dedupe_place; -recover=false; -help=false; -verbose=false; -quiet=false)
* Log and perform action

'-help' route (default if '-path' and '-recover' are unset)
* Output help message and exit

'-path' route (outputs overall progress semi-regularly by default; only errors with '-quiet'; every action if '-verbose'):
* Run 'fdupes -nr $DEDUPE_PATH' to get file lists ([][]string, with []string's all being identical; consecutive newlines in fdupes' output delimit different contents)
* Filter files in a group to only standard files and record metadata to preserve
* Log planned actions in case of crash
* Flush filesystem buffers really thoroughly to make sure the log isn't lost
* For each file group:
** For each file in group after first:
*** If file is owned by a different user and the program isn't being ran by root, then
**** Skip file and warn that it can't be properly recreated without root
**** Continue with next file
*** Delete file
*** Use 'cp --reflink' to make CoW copy
*** Reset metadata of capy to match original
*** Log that the file group has been processed
* Sync logs

'-recover' route (outputs recovery status and decision by default; errors only if '-quiet'; every bit of found info if '-verbose'):
* Check for 'plan' and 'action' logs, noting for each if it's 'valid', 'invalid', or 'absent'
* If both logs are 'absent', then ('nothing to do')
** Say that logs weren't found and exit
* If 'plan' log is not 'valid' and there is at least one log found, then ('can't fix this')
** Warn that the found log status is invalid, tell user to manually inspect logs at $LOCATION, and exit
* If 'plan' and 'action' logs are both 'valid' and match perfectly, then ('looks good; let's make sure')
** Say that the logs look like everything's fine, but we're checking to be safe.
** Go through each file listed in the 'plan' log and check that the metadata matches.
** Report any discrepencies.
* If 'plan' log is 'valid' and 'action' log is not 'valid' or doesn't match perfectly, then ('redo everything')
** Say that deduplication appears to have been interrupted in the middle and that we're redoing the plan.
** Jump into '-path' route after the 'plan' part.

*/
