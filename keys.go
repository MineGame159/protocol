// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

const (
	keyLine                              = "line"
	keyCharacter                         = "character"
	keyStart                             = "start"
	keyEnd                               = "end"
	keyURI                               = "uri"
	keyRange                             = "range"
	keyOriginSelectionRange              = "originSelectionRange"
	keyTargetURI                         = "targetUri"
	keyTargetRange                       = "targetRange"
	keyTargetSelectionRange              = "targetSelectionRange"
	keySeverity                          = "severity"
	keyCode                              = "code"
	keySource                            = "source"
	keyMessage                           = "message"
	keyRelatedInformation                = "relatedInformation"
	keyLocation                          = "location"
	keyTitle                             = "title"
	keyCommand                           = "command"
	keyArguments                         = "arguments"
	keyNewText                           = "newText"
	keyTextDocument                      = "textDocument"
	keyEdits                             = "edits"
	keyOverwrite                         = "overwrite"
	keyIgnoreIfExists                    = "ignoreIfExists"
	keyKind                              = "kind"
	keyOptions                           = "options"
	keyOldURI                            = "oldUri"
	keyNewURI                            = "newUri"
	keyRecursive                         = "recursive"
	keyIgnoreIfNotExists                 = "ignoreIfNotExists"
	keyChanges                           = "changes"
	keyDocumentChanges                   = "documentChanges"
	keyLanguageID                        = "languageId"
	keyVersion                           = "version"
	keyText                              = "text"
	keyPosition                          = "position"
	keyLanguage                          = "language"
	keyScheme                            = "scheme"
	keyPattern                           = "pattern"
	keyValue                             = "value"
	keyDiagnostics                       = "diagnostics"
	keyProcessID                         = "processId"
	keyRootPath                          = "rootPath"
	keyRootURI                           = "rootUri"
	keyInitializationOptions             = "initializationOptions"
	keyCapabilities                      = "capabilities"
	keyTrace                             = "trace"
	keyWorkspaceFolders                  = "workspaceFolders"
	keyFailureHandling                   = "failureHandling"
	keyResourceOperations                = "resourceOperations"
	keyDynamicRegistration               = "dynamicRegistration"
	keySymbolKind                        = "symbolKind"
	keyValueSet                          = "valueSet"
	keyApplyEdit                         = "applyEdit"
	keyWorkspaceEdit                     = "workspaceEdit"
	keyDidChangeConfiguration            = "didChangeConfiguration"
	keyDidChangeWatchedFiles             = "didChangeWatchedFiles"
	keySymbol                            = "symbol"
	keyExecuteCommand                    = "executeCommand"
	keyConfiguration                     = "configuration"
	keyDidSave                           = "didSave"
	keyWillSave                          = "willSave"
	keyWillSaveWaitUntil                 = "willSaveWaitUntil"
	keyCompletionItem                    = "completionItem"
	keyCompletionItemKind                = "completionItemKind"
	keyContextSupport                    = "contextSupport"
	keySnippetSupport                    = "snippetSupport"
	keyCommitCharactersSupport           = "commitCharactersSupport"
	keyDocumentationFormat               = "documentationFormat"
	keyDeprecatedSupport                 = "deprecatedSupport"
	keyPreselectSupport                  = "preselectSupport"
	keyContentFormat                     = "contentFormat"
	keySignatureInformation              = "signatureInformation"
	keyHierarchicalDocumentSymbolSupport = "hierarchicalDocumentSymbolSupport"
	keyLinkSupport                       = "linkSupport"
	keyCodeActionLiteralSupport          = "codeActionLiteralSupport"
	keyCodeActionKind                    = "codeActionKind"
	keyPrepareSupport                    = "prepareSupport"
	keyRangeLimit                        = "rangeLimit"
	keyLineFoldingOnly                   = "lineFoldingOnly"
	keySynchronization                   = "synchronization"
	keyCompletion                        = "completion"
	keyHover                             = "hover"
	keySignatureHelp                     = "signatureHelp"
	keyReferences                        = "references"
	keyDocumentHighlight                 = "documentHighlight"
	keyDocumentSymbol                    = "documentSymbol"
	keyFormatting                        = "formatting"
	keyRangeFormatting                   = "rangeFormatting"
	keyOnTypeFormatting                  = "onTypeFormatting"
	keyDeclaration                       = "declaration"
	keyDefinition                        = "definition"
	keyTypeDefinition                    = "typeDefinition"
	keyImplementation                    = "implementation"
	keyCodeAction                        = "codeAction"
	keyCodeLens                          = "codeLens"
	keyDocumentLink                      = "documentLink"
	keyColorProvider                     = "colorProvider"
	keyRename                            = "rename"
	keyPublishDiagnostics                = "publishDiagnostics"
	keyFoldingRange                      = "foldingRange"
	keyWorkspace                         = "workspace"
	keyExperimental                      = "experimental"
	keyRetry                             = "retry"
	keyResolveProvider                   = "resolveProvider"
	keyTriggerCharacters                 = "triggerCharacters"
	keyCodeActionKinds                   = "codeActionKinds"
	keyFirstTriggerCharacter             = "firstTriggerCharacter"
	keyMoreTriggerCharacter              = "moreTriggerCharacter"
	keyPrepareProvider                   = "prepareProvider"
	keyCommands                          = "commands"
	keyIncludeText                       = "includeText"
	keyOpenClose                         = "openClose"
	keyChange                            = "change"
	keySave                              = "save"
	keyID                                = "id"
	keySupported                         = "supported"
	keyChangeNotifications               = "changeNotifications"
	keyTextDocumentSync                  = "textDocumentSync"
	keyHoverProvider                     = "hoverProvider"
	keyCompletionProvider                = "completionProvider"
	keySignatureHelpProvider             = "signatureHelpProvider"
	keyDefinitionProvider                = "definitionProvider"
	keyTypeDefinitionProvider            = "typeDefinitionProvider"
	keyImplementationProvider            = "implementationProvider"
	keyReferencesProvider                = "referencesProvider"
	keyDocumentHighlightProvider         = "documentHighlightProvider"
	keyDocumentSymbolProvider            = "documentSymbolProvider"
	keyWorkspaceSymbolProvider           = "workspaceSymbolProvider"
	keyCodeActionProvider                = "codeActionProvider"
	keyCodeLensProvider                  = "codeLensProvider"
	keyDocumentFormattingProvider        = "documentFormattingProvider"
	keyDocumentRangeFormattingProvider   = "documentRangeFormattingProvider"
	keyDocumentOnTypeFormattingProvider  = "documentOnTypeFormattingProvider"
	keyRenameProvider                    = "renameProvider"
	keyDocumentLinkProvider              = "documentLinkProvider"
	keyFoldingRangeProvider              = "foldingRangeProvider"
	keyExecuteCommandProvider            = "executeCommandProvider"
	keyDocumentSelector                  = "documentSelector"
	keyContext                           = "context"
	keyTriggerCharacter                  = "triggerCharacter"
	keyTriggerKind                       = "triggerKind"
	keyIsIncomplete                      = "isIncomplete"
	keyItems                             = "items"
	keyAdditionalTextEdits               = "additionalTextEdits"
	keyCommitCharacters                  = "commitCharacters"
	keyData                              = "data"
	keyDeprecated                        = "deprecated"
	keyDetail                            = "detail"
	keyDocumentation                     = "documentation"
	keyFilterText                        = "filterText"
	keyInsertText                        = "insertText"
	keyInsertTextFormat                  = "insertTextFormat"
	keyLabel                             = "label"
	keyPreselect                         = "preselect"
	keySortText                          = "sortText"
	keyTextEdit                          = "textEdit"
	keyContents                          = "contents"
	keySignatures                        = "signatures"
	keyActiveParameter                   = "activeParameter"
	keyActiveSignature                   = "activeSignature"
	keyParameterInformation              = "parameterInformation"
	keyIncludeDeclaration                = "includeDeclaration"
	keyName                              = "name"
	keySelectionRange                    = "selectionRange"
	keyChildren                          = "children"
	keyContainerName                     = "containerName"
	keyOnly                              = "only"
	keyEdit                              = "edit"
	keyTarget                            = "target"
	keyColor                             = "color"
	keyAlpha                             = "alpha"
	keyBlue                              = "blue"
	keyGreen                             = "green"
	keyRed                               = "red"
	keyInsertSpaces                      = "insertSpaces"
	keyTabSize                           = "tabSize"
	keyCh                                = "ch"
	keyNewName                           = "newName"
	keyStartLine                         = "startLine"
	keyStartCharacter                    = "startCharacter"
	keyEndLine                           = "endLine"
	keyEndCharacter                      = "endCharacter"
	keyMethod                            = "method"
	keyRegisterOptions                   = "registerOptions"
	keyRegistrations                     = "registrations"
	keyUnregisterations                  = "unregisterations"
	keyContentChanges                    = "contentChanges"
	keyRangeLength                       = "rangeLength"
	keySyncKind                          = "syncKind"
	keyReason                            = "reason"
	keyType                              = "type"
	keyActions                           = "actions"
	keyEvent                             = "event"
	keyAdded                             = "added"
	keyRemoved                           = "removed"
	keySettings                          = "settings"
	keyScopeURI                          = "scopeUri"
	keySection                           = "section"
	keyWatchers                          = "watchers"
	keyGlobPattern                       = "globPattern"
	keyQuery                             = "query"
	keyApplied                           = "applied"
)
