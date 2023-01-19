// <<generate>>

package core

// Source: <main/php_compat.h>

/*
  +----------------------------------------------------------------------+
  | PHP Version 7                                                        |
  +----------------------------------------------------------------------+
  | Copyright (c) The PHP Group                                          |
  +----------------------------------------------------------------------+
  | This source file is subject to version 3.01 of the PHP license,      |
  | that is bundled with this package in the file LICENSE, and is        |
  | available through the world-wide-web at the following url:           |
  | http://www.php.net/license/3_01.txt                                  |
  | If you did not receive a copy of the PHP license and are unable to   |
  | obtain it through the world-wide-web, please send a note to          |
  | license@php.net so we can mail you a copy immediately.               |
  +----------------------------------------------------------------------+
  | Author:                                                              |
  +----------------------------------------------------------------------+
*/

// #define PHP_COMPAT_H

// # include < php_config . h >

// #define pcre2_jit_callback_8       php_pcre2_jit_callback

// #define pcre2_callout_enumerate_8       php_pcre2_callout_enumerate

// #define pcre2_code_copy_8       php_pcre2_code_copy

// #define pcre2_code_copy_with_tables_8       php_pcre2_code_copy_with_tables

// #define pcre2_code_free_8       php_pcre2_code_free

// #define pcre2_compile_8       php_pcre2_compile

// #define pcre2_compile_context_copy_8       php_pcre2_compile_context_copy

// #define pcre2_compile_context_create_8       php_pcre2_compile_context_create

// #define pcre2_compile_context_free_8       php_pcre2_compile_context_free

// #define pcre2_config_8       php_pcre2_config

// #define pcre2_convert_context_copy_8       php_pcre2_convert_context_copy

// #define pcre2_convert_context_create_8       php_pcre2_convert_context_create

// #define pcre2_convert_context_free_8       php_pcre2_convert_context_free

// #define pcre2_dfa_match_8       php_pcre2_dfa_match

// #define pcre2_general_context_copy_8       php_pcre2_general_context_copy

// #define pcre2_general_context_create_8       php_pcre2_general_context_create

// #define pcre2_general_context_free_8       php_pcre2_general_context_free

// #define pcre2_get_error_message_8       php_pcre2_get_error_message

// #define pcre2_get_mark_8       php_pcre2_get_mark

// #define pcre2_get_ovector_pointer_8       php_pcre2_get_ovector_pointer

// #define pcre2_get_ovector_count_8       php_pcre2_get_ovector_count

// #define pcre2_get_startchar_8       php_pcre2_get_startchar

// #define pcre2_jit_compile_8       php_pcre2_jit_compile

// #define pcre2_jit_match_8       php_pcre2_jit_match

// #define pcre2_jit_free_unused_memory_8       php_pcre2_jit_free_unused_memory

// #define pcre2_jit_stack_assign_8       php_pcre2_jit_stack_assign

// #define pcre2_jit_stack_create_8       php_pcre2_jit_stack_create

// #define pcre2_jit_stack_free_8       php_pcre2_jit_stack_free

// #define pcre2_maketables_8       php_pcre2_maketables

// #define pcre2_match_8       php_pcre2_match

// #define pcre2_match_context_copy_8       php_pcre2_match_context_copy

// #define pcre2_match_context_create_8       php_pcre2_match_context_create

// #define pcre2_match_context_free_8       php_pcre2_match_context_free

// #define pcre2_match_data_create_8       php_pcre2_match_data_create

// #define pcre2_match_data_create_from_pattern_8       php_pcre2_match_data_create_from_pattern

// #define pcre2_match_data_free_8       php_pcre2_match_data_free

// #define pcre2_pattern_info_8       php_pcre2_pattern_info

// #define pcre2_serialize_decode_8       php_pcre2_serialize_decode

// #define pcre2_serialize_encode_8       php_pcre2_serialize_encode

// #define pcre2_serialize_free_8       php_pcre2_serialize_free

// #define pcre2_serialize_get_number_of_codes_8       php_pcre2_serialize_get_number_of_codes

// #define pcre2_set_bsr_8       php_pcre2_set_bsr

// #define pcre2_set_callout_8       php_pcre2_set_callout

// #define pcre2_set_character_tables_8       php_pcre2_set_character_tables

// #define pcre2_set_compile_extra_options_8       php_pcre2_set_compile_extra_options

// #define pcre2_set_compile_recursion_guard_8       php_pcre2_set_compile_recursion_guard

// #define pcre2_set_depth_limit_8       php_pcre2_set_depth_limit

// #define pcre2_set_glob_escape_8       php_pcre2_set_glob_escape

// #define pcre2_set_glob_separator_8       php_pcre2_set_glob_separator

// #define pcre2_set_heap_limit_8       php_pcre2_set_heap_limit

// #define pcre2_set_match_limit_8       php_pcre2_set_match_limit

// #define pcre2_set_max_pattern_length_8       php_pcre2_set_max_pattern_length

// #define pcre2_set_newline_8       php_pcre2_set_newline

// #define pcre2_set_parens_nest_limit_8       php_pcre2_set_parens_nest_limit

// #define pcre2_set_offset_limit_8       php_pcre2_set_offset_limit

// #define pcre2_substitute_8       php_pcre2_substitute

// #define pcre2_substring_copy_byname_8       php_pcre2_substring_copy_byname

// #define pcre2_substring_copy_bynumber_8       php_pcre2_substring_copy_bynumber

// #define pcre2_substring_free_8       php_pcre2_substring_free

// #define pcre2_substring_get_byname_8       php_pcre2_substring_get_byname

// #define pcre2_substring_get_bynumber_8       php_pcre2_substring_get_bynumber

// #define pcre2_substring_length_byname_8       php_pcre2_substring_length_byname

// #define pcre2_substring_length_bynumber_8       php_pcre2_substring_length_bynumber

// #define pcre2_substring_list_get_8       php_pcre2_substring_list_get

// #define pcre2_substring_list_free_8       php_pcre2_substring_list_free

// #define pcre2_substring_nametable_scan_8       php_pcre2_substring_nametable_scan

// #define pcre2_substring_number_from_name_8       php_pcre2_substring_number_from_name

// #define pcre2_set_recursion_limit_8       php_pcre2_set_recursion_limit

// #define pcre2_set_recursion_memory_management_8       php_pcre2_set_recursion_memory_management

// #define lookup       php_lookup

// #define hashTableInit       php_hashTableInit

// #define hashTableDestroy       php_hashTableDestroy

// #define hashTableIterInit       php_hashTableIterInit

// #define hashTableIterNext       php_hashTableIterNext

// #define XML_DefaultCurrent       php_XML_DefaultCurrent

// #define XML_ErrorString       php_XML_ErrorString

// #define XML_ExpatVersion       php_XML_ExpatVersion

// #define XML_ExpatVersionInfo       php_XML_ExpatVersionInfo

// #define XML_ExternalEntityParserCreate       php_XML_ExternalEntityParserCreate

// #define XML_GetBase       php_XML_GetBase

// #define XML_GetBuffer       php_XML_GetBuffer

// #define XML_GetCurrentByteCount       php_XML_GetCurrentByteCount

// #define XML_GetCurrentByteIndex       php_XML_GetCurrentByteIndex

// #define XML_GetCurrentColumnNumber       php_XML_GetCurrentColumnNumber

// #define XML_GetCurrentLineNumber       php_XML_GetCurrentLineNumber

// #define XML_GetErrorCode       php_XML_GetErrorCode

// #define XML_GetIdAttributeIndex       php_XML_GetIdAttributeIndex

// #define XML_GetInputContext       php_XML_GetInputContext

// #define XML_GetSpecifiedAttributeCount       php_XML_GetSpecifiedAttributeCount

// #define XmlGetUtf16InternalEncodingNS       php_XmlGetUtf16InternalEncodingNS

// #define XmlGetUtf16InternalEncoding       php_XmlGetUtf16InternalEncoding

// #define XmlGetUtf8InternalEncodingNS       php_XmlGetUtf8InternalEncodingNS

// #define XmlGetUtf8InternalEncoding       php_XmlGetUtf8InternalEncoding

// #define XmlInitEncoding       php_XmlInitEncoding

// #define XmlInitEncodingNS       php_XmlInitEncodingNS

// #define XmlInitUnknownEncoding       php_XmlInitUnknownEncoding

// #define XmlInitUnknownEncodingNS       php_XmlInitUnknownEncodingNS

// #define XML_ParseBuffer       php_XML_ParseBuffer

// #define XML_Parse       php_XML_Parse

// #define XML_ParserCreate_MM       php_XML_ParserCreate_MM

// #define XML_ParserCreateNS       php_XML_ParserCreateNS

// #define XML_ParserCreate       php_XML_ParserCreate

// #define XML_ParserFree       php_XML_ParserFree

// #define XmlParseXmlDecl       php_XmlParseXmlDecl

// #define XmlParseXmlDeclNS       php_XmlParseXmlDeclNS

// #define XmlPrologStateInitExternalEntity       php_XmlPrologStateInitExternalEntity

// #define XmlPrologStateInit       php_XmlPrologStateInit

// #define XML_SetAttlistDeclHandler       php_XML_SetAttlistDeclHandler

// #define XML_SetBase       php_XML_SetBase

// #define XML_SetCdataSectionHandler       php_XML_SetCdataSectionHandler

// #define XML_SetCharacterDataHandler       php_XML_SetCharacterDataHandler

// #define XML_SetCommentHandler       php_XML_SetCommentHandler

// #define XML_SetDefaultHandlerExpand       php_XML_SetDefaultHandlerExpand

// #define XML_SetDefaultHandler       php_XML_SetDefaultHandler

// #define XML_SetDoctypeDeclHandler       php_XML_SetDoctypeDeclHandler

// #define XML_SetElementDeclHandler       php_XML_SetElementDeclHandler

// #define XML_SetElementHandler       php_XML_SetElementHandler

// #define XML_SetEncoding       php_XML_SetEncoding

// #define XML_SetEndCdataSectionHandler       php_XML_SetEndCdataSectionHandler

// #define XML_SetEndDoctypeDeclHandler       php_XML_SetEndDoctypeDeclHandler

// #define XML_SetEndElementHandler       php_XML_SetEndElementHandler

// #define XML_SetEndNamespaceDeclHandler       php_XML_SetEndNamespaceDeclHandler

// #define XML_SetEntityDeclHandler       php_XML_SetEntityDeclHandler

// #define XML_SetExternalEntityRefHandlerArg       php_XML_SetExternalEntityRefHandlerArg

// #define XML_SetExternalEntityRefHandler       php_XML_SetExternalEntityRefHandler

// #define XML_SetNamespaceDeclHandler       php_XML_SetNamespaceDeclHandler

// #define XML_SetNotationDeclHandler       php_XML_SetNotationDeclHandler

// #define XML_SetNotStandaloneHandler       php_XML_SetNotStandaloneHandler

// #define XML_SetParamEntityParsing       php_XML_SetParamEntityParsing

// #define XML_SetProcessingInstructionHandler       php_XML_SetProcessingInstructionHandler

// #define XML_SetReturnNSTriplet       php_XML_SetReturnNSTriplet

// #define XML_SetStartCdataSectionHandler       php_XML_SetStartCdataSectionHandler

// #define XML_SetStartDoctypeDeclHandler       php_XML_SetStartDoctypeDeclHandler

// #define XML_SetStartElementHandler       php_XML_SetStartElementHandler

// #define XML_SetStartNamespaceDeclHandler       php_XML_SetStartNamespaceDeclHandler

// #define XML_SetUnknownEncodingHandler       php_XML_SetUnknownEncodingHandler

// #define XML_SetUnparsedEntityDeclHandler       php_XML_SetUnparsedEntityDeclHandler

// #define XML_SetUserData       php_XML_SetUserData

// #define XML_SetXmlDeclHandler       php_XML_SetXmlDeclHandler

// #define XmlSizeOfUnknownEncoding       php_XmlSizeOfUnknownEncoding

// #define XML_UseParserAsHandlerArg       php_XML_UseParserAsHandlerArg

// #define XmlUtf16Encode       php_XmlUtf16Encode

// #define XmlUtf8Encode       php_XmlUtf8Encode

// #define XML_FreeContentModel       php_XML_FreeContentModel

// #define XML_MemMalloc       php_XML_MemMalloc

// #define XML_MemRealloc       php_XML_MemRealloc

// #define XML_MemFree       php_XML_MemFree

// #define XML_UseForeignDTD       php_XML_UseForeignDTD

// #define XML_GetFeatureList       php_XML_GetFeatureList

// #define XML_ParserReset       php_XML_ParserReset

/* Define to specify how much context to retain around the current parse
   point. */

// #define XML_CONTEXT_BYTES       1024

/* Define to make parameter entity parsing functionality available. */

// #define XML_DTD       1

/* Define to make XML Namespaces functionality available. */

// #define XML_NS       1
