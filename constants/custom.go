package constants

// StringRepoAlreadyExists is a custom error message
const StringRepoAlreadyExists string = "Unable to clone repository: Another git repository already exists in target location. "

// StringNoExternalGitnoArguments is a custom error message
const StringNoExternalGitnoArguments string = "External 'git' not found, can not pass any additional arguments to 'git clone'. "

// StringLinkDoesNotExist is a custom error message
const StringLinkDoesNotExist string = "Unable to link repository: Can not open source repository. "

// StringLinkAlreadyExists is a custom error message
const StringLinkAlreadyExists string = "Unable to link repository: Another directory already exists in target location. "

// StringLinkSamePath is a custom error message
const StringLinkSamePath string = "Unable to link repository: Link source and target are identical. "
