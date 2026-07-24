# GitHub commit attribution: how it works, how it broke here, how to fix it

This is a runbook for the problem where a commit shows up on GitHub under the
wrong person (or as a faded/unlinked name) even though `git log` locally looks
fine. It documents the exact incident on this repo and the general mechanism, so
it can be diagnosed and fixed again in seconds.

## TL;DR of the incident

The initial commit of `Brilhante29/kiri-aws` was attributed to a different
account. Cause: the commit's **committer email** was
`guilhermebrilhante@users.noreply.github.com`, which is the no-reply address of a
*different* GitHub account (`guilhermebrilhante`), not `Brilhante29`. GitHub maps
commits to accounts by email, so it linked the commit to the account that owns
that address.

Fix applied: rewrote the root commit's author and committer to
`Guilherme Brilhante <85261647+Brilhante29@users.noreply.github.com>` (the
ID-based no-reply of the `Brilhante29` account, id `85261647`) and force-pushed.

```bash
git config user.name  "Guilherme Brilhante"
git config user.email "85261647+Brilhante29@users.noreply.github.com"
GIT_COMMITTER_NAME="Guilherme Brilhante" \
GIT_COMMITTER_EMAIL="85261647+Brilhante29@users.noreply.github.com" \
  git commit --amend --reset-author --no-edit
git push --force-with-lease origin main
```

Result: `0eda0ec...1b1ec26 main -> main (forced update)`.

## The mechanism

A git commit carries two identities, each a name plus an email:

- **author**: who wrote the change (`%an <%ae>`).
- **committer**: who created the commit object (`%cn <%ce>`). Amend, rebase, and
  cherry-pick set this to whoever ran the command.

GitHub does not care about the *name*. It attributes a commit to a user account
by matching the **email** against:

1. a **verified** email on some account's Settings → Emails, or
2. that account's **no-reply** address.

Whichever account owns the matching email gets the avatar and the link. The name
string in the commit is only shown as a label. So two things can go wrong:

- The email belongs to **another** account (this incident).
- The email belongs to **no** account, so GitHub shows a plain, unlinked name
  with a default avatar.

Both author and committer are displayed, and the repository's commit list uses
the **committer** for the "committed" line, which is why a wrong committer email
alone is enough to misattribute.

## No-reply email formats

GitHub issues a privacy no-reply address per account. Two formats exist:

- Modern, ID-based: `ID+username@users.noreply.github.com`
  (e.g. `85261647+Brilhante29@users.noreply.github.com`). Always maps to the
  account with that numeric id, regardless of username changes.
- Legacy, name-based: `username@users.noreply.github.com`
  (e.g. `guilhermebrilhante@users.noreply.github.com`). Maps by username only,
  and follows whoever holds that username.

The incident address `guilhermebrilhante@users.noreply.github.com` is the legacy
form for username `guilhermebrilhante`, a separate account from `Brilhante29`.
Prefer the ID-based form: it is unambiguous and survives username changes.

## Diagnosis

Show the real author and committer of every commit (names alone lie, look at
emails):

```bash
git log --format='%h  A:%an <%ae>  C:%cn <%ce>'
```

Find the numeric id and canonical no-reply of the account you *want* the commit
attributed to (public, no auth needed):

```bash
curl -s https://api.github.com/users/Brilhante29
# read .id and .login  ->  <id>+<login>@users.noreply.github.com
```

Check what git will stamp on the next commit:

```bash
git config user.name;  git config user.email          # repo-local
git config --global user.name;  git config --global user.email
```

If the global email is the wrong no-reply, every repo inherits the bug until the
local config overrides it.

## Fix by number of commits

### Single commit (root or HEAD)

As in the incident:

```bash
git config user.email "<id>+<login>@users.noreply.github.com"
GIT_COMMITTER_EMAIL="<id>+<login>@users.noreply.github.com" \
  git commit --amend --reset-author --no-edit
git push --force-with-lease origin <branch>
```

`--reset-author` rewrites the **author** to the current config identity; the
**committer** is always taken from config/env on any commit operation, so both
end up correct.

### Whole history (many commits)

Use `git filter-repo` (install: `pip install git-filter-repo`). Rewrite only the
commits that carry the wrong email:

```bash
git filter-repo --force --email-callback '
  return b"85261647+Brilhante29@users.noreply.github.com"
    if email in (b"guilhermebrilhante@users.noreply.github.com",
                 b"guilhermebrilhante00@gmail.com")
    else email
'
# filter-repo drops the remote; re-add and force-push
git remote add origin https://github.com/Brilhante29/kiri-aws.git
git push --force-with-lease origin --all
```

If `filter-repo` is unavailable, `git rebase --root -x 'git commit --amend
--reset-author --no-edit'` works for linear history but is slower and touches
every commit.

## Force-push safety

- Rewriting a published commit changes its hash, so the remote must be
  overwritten. Use `--force-with-lease`, not `--force`: it refuses the push if the
  remote moved since your last fetch, protecting a teammate's work.
- Anyone who already cloned keeps the old hashes. On a solo repo this is a
  non-issue; on a shared repo, coordinate first.
- Open pull requests built on the old commits may need rebasing.

## Prevention

1. Set the right email once, globally, using the ID-based no-reply:

   ```bash
   git config --global user.name  "Guilherme Brilhante"
   git config --global user.email "85261647+Brilhante29@users.noreply.github.com"
   ```

2. On GitHub, Settings → Emails → enable "Keep my email addresses private" and
   "Block command line pushes that expose my email". GitHub then shows the exact
   ID-based no-reply to use and rejects pushes that leak a real address.

3. If you prefer a real email as author, add it under Settings → Emails and
   **verify** it. Only verified emails attribute; an unverified one shows unlinked.

4. Watch out for tools that set git identity behind your back: `gh auth login`,
   IDE Git integrations, and CI checkout actions can inject a no-reply for the
   account they authenticated as. After such a login, re-check
   `git config user.email` in the repo.

5. For a repo shared by several identities, commit a `.mailmap` so tools and
   `git shortlog` collapse them to one canonical name/email without rewriting
   history:

   ```
   Guilherme Brilhante <85261647+Brilhante29@users.noreply.github.com> <guilhermebrilhante00@gmail.com>
   ```

## Verify the fix landed

```bash
git log -1 --format='A:%an <%ae>%nC:%cn <%ce>'
# both lines must read: Guilherme Brilhante <85261647+Brilhante29@users.noreply.github.com>
```

Then reload the commit on GitHub. Attribution updates as soon as the pushed
email is recognized; the avatar may take a moment to refresh.
