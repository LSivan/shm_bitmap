# shm_bitmap
Bitmap based on shared memory.

## Why bitmap

reference: [Bloom_filter](https://en.wikipedia.org/wiki/Bloom_filter)

## Why shared memory

For crash-safe. 

In the actual production case, it may take lots of time when application construct a large bitmap.

If shared memory is not used, it means the application cannot function properly upon recovery from a crash.   