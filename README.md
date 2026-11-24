# Daily Temperatures

> NeetCode 150 · Stack — next-greater-element reported as a distance: resolve stacked colder days the moment a warmer one arrives. Go, stdlib only, tested.

## Problem

For each day's temperature, how many days until a strictly warmer one (0 if it never comes)? `[73,74,75,71,69,72,76,73]` → `[1,1,4,2,1,1,0,0]`.

## Approach

This is dsa-080's next-greater-element wearing a weather report — the answer is the *index gap* instead of the value. Keep a stack of day indexes whose warmer day hasn't arrived; each new temperature pops and resolves every colder stacked day (`out[j] = i − j`), then waits its own turn. Strictness matters: equal temperatures don't resolve (pinned by the all-70s test), so the pop condition is `<`, keeping equal runs stacked. Push-once-pop-once gives linearity; unresolved survivors keep their zero from the output's initialization.

**Complexity:** O(n) time, O(n) stack.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
