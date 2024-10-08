// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package mutagenbox

// mutagenbox is a package that encapsulates state and logic specific to how
// we need to manage mutagen for the codex cloud.
//
// Also, resolves some compile cycles:
//   - [cloud] depends on [mutagenbox], [sshshim], and [mutagen].
//   - [sshshim] depends on [mutagenbox] and [mutagen].
