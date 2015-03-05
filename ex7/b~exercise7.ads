pragma Ada_95;
with System;
package ada_main is
   pragma Warnings (Off);

   gnat_argc : Integer;
   gnat_argv : System.Address;
   gnat_envp : System.Address;

   pragma Import (C, gnat_argc);
   pragma Import (C, gnat_argv);
   pragma Import (C, gnat_envp);

   gnat_exit_status : Integer;
   pragma Import (C, gnat_exit_status);

   GNAT_Version : constant String :=
                    "GNAT Version: 4.6" & ASCII.NUL;
   pragma Export (C, GNAT_Version, "__gnat_version");

   Ada_Main_Program_Name : constant String := "_ada_exercise7" & ASCII.NUL;
   pragma Export (C, Ada_Main_Program_Name, "__gnat_ada_main_program_name");

   procedure adafinal;
   pragma Export (C, adafinal, "adafinal");

   procedure adainit;
   pragma Export (C, adainit, "adainit");

   procedure Break_Start;
   pragma Import (C, Break_Start, "__gnat_break_start");

   function main
     (argc : Integer;
      argv : System.Address;
      envp : System.Address)
      return Integer;
   pragma Export (C, main, "main");

   type Version_32 is mod 2 ** 32;
   u00001 : constant Version_32 := 16#420c9368#;
   pragma Export (C, u00001, "exercise7B");
   u00002 : constant Version_32 := 16#ba46b2cd#;
   pragma Export (C, u00002, "system__standard_libraryB");
   u00003 : constant Version_32 := 16#1e2e640d#;
   pragma Export (C, u00003, "system__standard_libraryS");
   u00004 : constant Version_32 := 16#3ffc8e18#;
   pragma Export (C, u00004, "adaS");
   u00005 : constant Version_32 := 16#df111b7b#;
   pragma Export (C, u00005, "ada__calendar__delaysB");
   u00006 : constant Version_32 := 16#ab1c3be0#;
   pragma Export (C, u00006, "ada__calendar__delaysS");
   u00007 : constant Version_32 := 16#0f244912#;
   pragma Export (C, u00007, "ada__calendarB");
   u00008 : constant Version_32 := 16#0bc00dc5#;
   pragma Export (C, u00008, "ada__calendarS");
   u00009 : constant Version_32 := 16#9229643d#;
   pragma Export (C, u00009, "ada__exceptionsB");
   u00010 : constant Version_32 := 16#e3df9d67#;
   pragma Export (C, u00010, "ada__exceptionsS");
   u00011 : constant Version_32 := 16#95643e9a#;
   pragma Export (C, u00011, "ada__exceptions__last_chance_handlerB");
   u00012 : constant Version_32 := 16#03cf4fc2#;
   pragma Export (C, u00012, "ada__exceptions__last_chance_handlerS");
   u00013 : constant Version_32 := 16#23e1f70b#;
   pragma Export (C, u00013, "systemS");
   u00014 : constant Version_32 := 16#30ec78bc#;
   pragma Export (C, u00014, "system__soft_linksB");
   u00015 : constant Version_32 := 16#e2ebe502#;
   pragma Export (C, u00015, "system__soft_linksS");
   u00016 : constant Version_32 := 16#0d2b82ae#;
   pragma Export (C, u00016, "system__parametersB");
   u00017 : constant Version_32 := 16#bfbc74f1#;
   pragma Export (C, u00017, "system__parametersS");
   u00018 : constant Version_32 := 16#72905399#;
   pragma Export (C, u00018, "system__secondary_stackB");
   u00019 : constant Version_32 := 16#378fd0a5#;
   pragma Export (C, u00019, "system__secondary_stackS");
   u00020 : constant Version_32 := 16#ace32e1e#;
   pragma Export (C, u00020, "system__storage_elementsB");
   u00021 : constant Version_32 := 16#d92c8a93#;
   pragma Export (C, u00021, "system__storage_elementsS");
   u00022 : constant Version_32 := 16#4f750b3b#;
   pragma Export (C, u00022, "system__stack_checkingB");
   u00023 : constant Version_32 := 16#80434b27#;
   pragma Export (C, u00023, "system__stack_checkingS");
   u00024 : constant Version_32 := 16#a7343537#;
   pragma Export (C, u00024, "system__exception_tableB");
   u00025 : constant Version_32 := 16#8120f83b#;
   pragma Export (C, u00025, "system__exception_tableS");
   u00026 : constant Version_32 := 16#ff3fa16b#;
   pragma Export (C, u00026, "system__htableB");
   u00027 : constant Version_32 := 16#cc3e5bd4#;
   pragma Export (C, u00027, "system__htableS");
   u00028 : constant Version_32 := 16#8b7dad61#;
   pragma Export (C, u00028, "system__string_hashB");
   u00029 : constant Version_32 := 16#057d2f9f#;
   pragma Export (C, u00029, "system__string_hashS");
   u00030 : constant Version_32 := 16#6a8a6a74#;
   pragma Export (C, u00030, "system__exceptionsB");
   u00031 : constant Version_32 := 16#86f01d0a#;
   pragma Export (C, u00031, "system__exceptionsS");
   u00032 : constant Version_32 := 16#b012ff50#;
   pragma Export (C, u00032, "system__img_intB");
   u00033 : constant Version_32 := 16#213a17c9#;
   pragma Export (C, u00033, "system__img_intS");
   u00034 : constant Version_32 := 16#dc8e33ed#;
   pragma Export (C, u00034, "system__tracebackB");
   u00035 : constant Version_32 := 16#4266237e#;
   pragma Export (C, u00035, "system__tracebackS");
   u00036 : constant Version_32 := 16#4900ab7d#;
   pragma Export (C, u00036, "system__unsigned_typesS");
   u00037 : constant Version_32 := 16#907d882f#;
   pragma Export (C, u00037, "system__wch_conB");
   u00038 : constant Version_32 := 16#9c0ad936#;
   pragma Export (C, u00038, "system__wch_conS");
   u00039 : constant Version_32 := 16#22fed88a#;
   pragma Export (C, u00039, "system__wch_stwB");
   u00040 : constant Version_32 := 16#b11bf537#;
   pragma Export (C, u00040, "system__wch_stwS");
   u00041 : constant Version_32 := 16#5d4d477e#;
   pragma Export (C, u00041, "system__wch_cnvB");
   u00042 : constant Version_32 := 16#82f45fe0#;
   pragma Export (C, u00042, "system__wch_cnvS");
   u00043 : constant Version_32 := 16#f77d8799#;
   pragma Export (C, u00043, "interfacesS");
   u00044 : constant Version_32 := 16#75729fba#;
   pragma Export (C, u00044, "system__wch_jisB");
   u00045 : constant Version_32 := 16#d686c4f4#;
   pragma Export (C, u00045, "system__wch_jisS");
   u00046 : constant Version_32 := 16#ada34a87#;
   pragma Export (C, u00046, "system__traceback_entriesB");
   u00047 : constant Version_32 := 16#71c0194a#;
   pragma Export (C, u00047, "system__traceback_entriesS");
   u00048 : constant Version_32 := 16#22d03640#;
   pragma Export (C, u00048, "system__os_primitivesB");
   u00049 : constant Version_32 := 16#93307b22#;
   pragma Export (C, u00049, "system__os_primitivesS");
   u00050 : constant Version_32 := 16#ee80728a#;
   pragma Export (C, u00050, "system__tracesB");
   u00051 : constant Version_32 := 16#d1fc9fa1#;
   pragma Export (C, u00051, "system__tracesS");
   u00052 : constant Version_32 := 16#f64b89a4#;
   pragma Export (C, u00052, "ada__integer_text_ioB");
   u00053 : constant Version_32 := 16#f1daf268#;
   pragma Export (C, u00053, "ada__integer_text_ioS");
   u00054 : constant Version_32 := 16#07116dec#;
   pragma Export (C, u00054, "ada__tagsB");
   u00055 : constant Version_32 := 16#21b957c3#;
   pragma Export (C, u00055, "ada__tagsS");
   u00056 : constant Version_32 := 16#68f8d5f8#;
   pragma Export (C, u00056, "system__val_lluB");
   u00057 : constant Version_32 := 16#33f2fc0f#;
   pragma Export (C, u00057, "system__val_lluS");
   u00058 : constant Version_32 := 16#46a1f7a9#;
   pragma Export (C, u00058, "system__val_utilB");
   u00059 : constant Version_32 := 16#284c6214#;
   pragma Export (C, u00059, "system__val_utilS");
   u00060 : constant Version_32 := 16#b7fa72e7#;
   pragma Export (C, u00060, "system__case_utilB");
   u00061 : constant Version_32 := 16#8efd9783#;
   pragma Export (C, u00061, "system__case_utilS");
   u00062 : constant Version_32 := 16#7a8f4ce5#;
   pragma Export (C, u00062, "ada__text_ioB");
   u00063 : constant Version_32 := 16#78993766#;
   pragma Export (C, u00063, "ada__text_ioS");
   u00064 : constant Version_32 := 16#1358602f#;
   pragma Export (C, u00064, "ada__streamsS");
   u00065 : constant Version_32 := 16#7a48d8b1#;
   pragma Export (C, u00065, "interfaces__c_streamsB");
   u00066 : constant Version_32 := 16#40dd1af2#;
   pragma Export (C, u00066, "interfaces__c_streamsS");
   u00067 : constant Version_32 := 16#13cbc5a8#;
   pragma Export (C, u00067, "system__crtlS");
   u00068 : constant Version_32 := 16#5efa797c#;
   pragma Export (C, u00068, "system__file_ioB");
   u00069 : constant Version_32 := 16#2e96f0e6#;
   pragma Export (C, u00069, "system__file_ioS");
   u00070 : constant Version_32 := 16#7cc77cc0#;
   pragma Export (C, u00070, "ada__finalizationB");
   u00071 : constant Version_32 := 16#01acb175#;
   pragma Export (C, u00071, "ada__finalizationS");
   u00072 : constant Version_32 := 16#01cb6d81#;
   pragma Export (C, u00072, "system__finalization_rootB");
   u00073 : constant Version_32 := 16#2d16f6f3#;
   pragma Export (C, u00073, "system__finalization_rootS");
   u00074 : constant Version_32 := 16#dbb36d26#;
   pragma Export (C, u00074, "system__finalization_implementationB");
   u00075 : constant Version_32 := 16#bdfa5ab4#;
   pragma Export (C, u00075, "system__finalization_implementationS");
   u00076 : constant Version_32 := 16#386436bc#;
   pragma Export (C, u00076, "system__restrictionsB");
   u00077 : constant Version_32 := 16#db039e46#;
   pragma Export (C, u00077, "system__restrictionsS");
   u00078 : constant Version_32 := 16#a6e358bc#;
   pragma Export (C, u00078, "system__stream_attributesB");
   u00079 : constant Version_32 := 16#e89b4b3f#;
   pragma Export (C, u00079, "system__stream_attributesS");
   u00080 : constant Version_32 := 16#b46168d5#;
   pragma Export (C, u00080, "ada__io_exceptionsS");
   u00081 : constant Version_32 := 16#a2230cb9#;
   pragma Export (C, u00081, "interfaces__cB");
   u00082 : constant Version_32 := 16#ebbc3ca7#;
   pragma Export (C, u00082, "interfaces__cS");
   u00083 : constant Version_32 := 16#7401caa7#;
   pragma Export (C, u00083, "interfaces__c__stringsB");
   u00084 : constant Version_32 := 16#739e0600#;
   pragma Export (C, u00084, "interfaces__c__stringsS");
   u00085 : constant Version_32 := 16#621b06ef#;
   pragma Export (C, u00085, "system__crtl__runtimeS");
   u00086 : constant Version_32 := 16#f74220e8#;
   pragma Export (C, u00086, "system__os_libB");
   u00087 : constant Version_32 := 16#a6d80a38#;
   pragma Export (C, u00087, "system__os_libS");
   u00088 : constant Version_32 := 16#4cd8aca0#;
   pragma Export (C, u00088, "system__stringsB");
   u00089 : constant Version_32 := 16#940bbdcf#;
   pragma Export (C, u00089, "system__stringsS");
   u00090 : constant Version_32 := 16#fcde1931#;
   pragma Export (C, u00090, "system__file_control_blockS");
   u00091 : constant Version_32 := 16#b90c86f6#;
   pragma Export (C, u00091, "ada__finalization__list_controllerB");
   u00092 : constant Version_32 := 16#b97dfd74#;
   pragma Export (C, u00092, "ada__finalization__list_controllerS");
   u00093 : constant Version_32 := 16#f6fdca1c#;
   pragma Export (C, u00093, "ada__text_io__integer_auxB");
   u00094 : constant Version_32 := 16#b9793d30#;
   pragma Export (C, u00094, "ada__text_io__integer_auxS");
   u00095 : constant Version_32 := 16#515dc0e3#;
   pragma Export (C, u00095, "ada__text_io__generic_auxB");
   u00096 : constant Version_32 := 16#a6c327d3#;
   pragma Export (C, u00096, "ada__text_io__generic_auxS");
   u00097 : constant Version_32 := 16#ef6c8032#;
   pragma Export (C, u00097, "system__img_biuB");
   u00098 : constant Version_32 := 16#8f222330#;
   pragma Export (C, u00098, "system__img_biuS");
   u00099 : constant Version_32 := 16#10618bf9#;
   pragma Export (C, u00099, "system__img_llbB");
   u00100 : constant Version_32 := 16#cee533ce#;
   pragma Export (C, u00100, "system__img_llbS");
   u00101 : constant Version_32 := 16#9777733a#;
   pragma Export (C, u00101, "system__img_lliB");
   u00102 : constant Version_32 := 16#32aea2da#;
   pragma Export (C, u00102, "system__img_lliS");
   u00103 : constant Version_32 := 16#f931f062#;
   pragma Export (C, u00103, "system__img_llwB");
   u00104 : constant Version_32 := 16#67891058#;
   pragma Export (C, u00104, "system__img_llwS");
   u00105 : constant Version_32 := 16#b532ff4e#;
   pragma Export (C, u00105, "system__img_wiuB");
   u00106 : constant Version_32 := 16#e163a4a2#;
   pragma Export (C, u00106, "system__img_wiuS");
   u00107 : constant Version_32 := 16#7993dbbd#;
   pragma Export (C, u00107, "system__val_intB");
   u00108 : constant Version_32 := 16#6b44dd34#;
   pragma Export (C, u00108, "system__val_intS");
   u00109 : constant Version_32 := 16#e6965fe6#;
   pragma Export (C, u00109, "system__val_unsB");
   u00110 : constant Version_32 := 16#59a84646#;
   pragma Export (C, u00110, "system__val_unsS");
   u00111 : constant Version_32 := 16#936e9286#;
   pragma Export (C, u00111, "system__val_lliB");
   u00112 : constant Version_32 := 16#b9c511ab#;
   pragma Export (C, u00112, "system__val_lliS");
   u00113 : constant Version_32 := 16#84ad4a42#;
   pragma Export (C, u00113, "ada__numericsS");
   u00114 : constant Version_32 := 16#ac5daf3d#;
   pragma Export (C, u00114, "ada__numerics__float_randomB");
   u00115 : constant Version_32 := 16#ac27f55b#;
   pragma Export (C, u00115, "ada__numerics__float_randomS");
   u00116 : constant Version_32 := 16#036cdb71#;
   pragma Export (C, u00116, "system__random_numbersB");
   u00117 : constant Version_32 := 16#8a2f4f98#;
   pragma Export (C, u00117, "system__random_numbersS");
   u00118 : constant Version_32 := 16#194ccd7b#;
   pragma Export (C, u00118, "system__img_unsB");
   u00119 : constant Version_32 := 16#d6f4978a#;
   pragma Export (C, u00119, "system__img_unsS");
   u00120 : constant Version_32 := 16#00dc945f#;
   pragma Export (C, u00120, "ada__real_timeB");
   u00121 : constant Version_32 := 16#5bfb6637#;
   pragma Export (C, u00121, "ada__real_timeS");
   u00122 : constant Version_32 := 16#763c483e#;
   pragma Export (C, u00122, "system__arith_64B");
   u00123 : constant Version_32 := 16#748dd1fb#;
   pragma Export (C, u00123, "system__arith_64S");
   u00124 : constant Version_32 := 16#a6fc8110#;
   pragma Export (C, u00124, "system__taskingB");
   u00125 : constant Version_32 := 16#61a61d3d#;
   pragma Export (C, u00125, "system__taskingS");
   u00126 : constant Version_32 := 16#e28875aa#;
   pragma Export (C, u00126, "system__task_primitivesS");
   u00127 : constant Version_32 := 16#c9494811#;
   pragma Export (C, u00127, "system__os_interfaceB");
   u00128 : constant Version_32 := 16#732a7471#;
   pragma Export (C, u00128, "system__os_interfaceS");
   u00129 : constant Version_32 := 16#6470edba#;
   pragma Export (C, u00129, "system__linuxS");
   u00130 : constant Version_32 := 16#52a12383#;
   pragma Export (C, u00130, "system__task_primitives__operationsB");
   u00131 : constant Version_32 := 16#e8297115#;
   pragma Export (C, u00131, "system__task_primitives__operationsS");
   u00132 : constant Version_32 := 16#b71e6964#;
   pragma Export (C, u00132, "system__bit_opsB");
   u00133 : constant Version_32 := 16#c30e4013#;
   pragma Export (C, u00133, "system__bit_opsS");
   u00134 : constant Version_32 := 16#2a8bb2a4#;
   pragma Export (C, u00134, "system__interrupt_managementB");
   u00135 : constant Version_32 := 16#8aeaffc0#;
   pragma Export (C, u00135, "system__interrupt_managementS");
   u00136 : constant Version_32 := 16#17ae9ccc#;
   pragma Export (C, u00136, "system__multiprocessorsB");
   u00137 : constant Version_32 := 16#1b4d6878#;
   pragma Export (C, u00137, "system__multiprocessorsS");
   u00138 : constant Version_32 := 16#cf99f883#;
   pragma Export (C, u00138, "system__stack_checking__operationsB");
   u00139 : constant Version_32 := 16#49df1cef#;
   pragma Export (C, u00139, "system__stack_checking__operationsS");
   u00140 : constant Version_32 := 16#3d54d5f6#;
   pragma Export (C, u00140, "system__task_infoB");
   u00141 : constant Version_32 := 16#6cfdb051#;
   pragma Export (C, u00141, "system__task_infoS");
   u00142 : constant Version_32 := 16#0a363cc0#;
   pragma Export (C, u00142, "system__tasking__debugB");
   u00143 : constant Version_32 := 16#71f06032#;
   pragma Export (C, u00143, "system__tasking__debugS");
   u00144 : constant Version_32 := 16#39591e91#;
   pragma Export (C, u00144, "system__concat_2B");
   u00145 : constant Version_32 := 16#d83105f7#;
   pragma Export (C, u00145, "system__concat_2S");
   u00146 : constant Version_32 := 16#ae97ef6c#;
   pragma Export (C, u00146, "system__concat_3B");
   u00147 : constant Version_32 := 16#55cbf561#;
   pragma Export (C, u00147, "system__concat_3S");
   u00148 : constant Version_32 := 16#c9fdc962#;
   pragma Export (C, u00148, "system__concat_6B");
   u00149 : constant Version_32 := 16#e42b021f#;
   pragma Export (C, u00149, "system__concat_6S");
   u00150 : constant Version_32 := 16#def1dd00#;
   pragma Export (C, u00150, "system__concat_5B");
   u00151 : constant Version_32 := 16#33d839aa#;
   pragma Export (C, u00151, "system__concat_5S");
   u00152 : constant Version_32 := 16#3493e6c0#;
   pragma Export (C, u00152, "system__concat_4B");
   u00153 : constant Version_32 := 16#21be14b5#;
   pragma Export (C, u00153, "system__concat_4S");
   u00154 : constant Version_32 := 16#1eab0e09#;
   pragma Export (C, u00154, "system__img_enum_newB");
   u00155 : constant Version_32 := 16#a4e63cfb#;
   pragma Export (C, u00155, "system__img_enum_newS");
   u00156 : constant Version_32 := 16#06417083#;
   pragma Export (C, u00156, "system__img_lluB");
   u00157 : constant Version_32 := 16#00c9abbe#;
   pragma Export (C, u00157, "system__img_lluS");
   u00158 : constant Version_32 := 16#3b607801#;
   pragma Export (C, u00158, "system__stack_usageB");
   u00159 : constant Version_32 := 16#abdf2d5b#;
   pragma Export (C, u00159, "system__stack_usageS");
   u00160 : constant Version_32 := 16#d7aac20c#;
   pragma Export (C, u00160, "system__ioB");
   u00161 : constant Version_32 := 16#bda30044#;
   pragma Export (C, u00161, "system__ioS");
   u00162 : constant Version_32 := 16#48087e79#;
   pragma Export (C, u00162, "system__tasking__protected_objectsB");
   u00163 : constant Version_32 := 16#e2c8a7eb#;
   pragma Export (C, u00163, "system__tasking__protected_objectsS");
   u00164 : constant Version_32 := 16#ff133197#;
   pragma Export (C, u00164, "system__soft_links__taskingB");
   u00165 : constant Version_32 := 16#da1c0a98#;
   pragma Export (C, u00165, "system__soft_links__taskingS");
   u00166 : constant Version_32 := 16#17d21067#;
   pragma Export (C, u00166, "ada__exceptions__is_null_occurrenceB");
   u00167 : constant Version_32 := 16#c4bf5e73#;
   pragma Export (C, u00167, "ada__exceptions__is_null_occurrenceS");
   u00168 : constant Version_32 := 16#7a9a5f5e#;
   pragma Export (C, u00168, "system__tasking__protected_objects__entriesB");
   u00169 : constant Version_32 := 16#a416dd28#;
   pragma Export (C, u00169, "system__tasking__protected_objects__entriesS");
   u00170 : constant Version_32 := 16#4d3e0469#;
   pragma Export (C, u00170, "system__tasking__initializationB");
   u00171 : constant Version_32 := 16#7f6b69f1#;
   pragma Export (C, u00171, "system__tasking__initializationS");
   u00172 : constant Version_32 := 16#9fcca227#;
   pragma Export (C, u00172, "system__tasking__protected_objects__operationsB");
   u00173 : constant Version_32 := 16#eb309cf0#;
   pragma Export (C, u00173, "system__tasking__protected_objects__operationsS");
   u00174 : constant Version_32 := 16#b59e6e78#;
   pragma Export (C, u00174, "system__tasking__entry_callsB");
   u00175 : constant Version_32 := 16#687efea4#;
   pragma Export (C, u00175, "system__tasking__entry_callsS");
   u00176 : constant Version_32 := 16#88081561#;
   pragma Export (C, u00176, "system__tasking__queuingB");
   u00177 : constant Version_32 := 16#ee1cad28#;
   pragma Export (C, u00177, "system__tasking__queuingS");
   u00178 : constant Version_32 := 16#f488ba38#;
   pragma Export (C, u00178, "system__tasking__utilitiesB");
   u00179 : constant Version_32 := 16#0ef27416#;
   pragma Export (C, u00179, "system__tasking__utilitiesS");
   u00180 : constant Version_32 := 16#bd6fc52e#;
   pragma Export (C, u00180, "system__traces__taskingB");
   u00181 : constant Version_32 := 16#becc801d#;
   pragma Export (C, u00181, "system__traces__taskingS");
   u00182 : constant Version_32 := 16#959cf892#;
   pragma Export (C, u00182, "system__tasking__rendezvousB");
   u00183 : constant Version_32 := 16#d056c48f#;
   pragma Export (C, u00183, "system__tasking__rendezvousS");
   u00184 : constant Version_32 := 16#14757b6f#;
   pragma Export (C, u00184, "system__tasking__stagesB");
   u00185 : constant Version_32 := 16#84f88b56#;
   pragma Export (C, u00185, "system__tasking__stagesS");
   u00186 : constant Version_32 := 16#57a37a42#;
   pragma Export (C, u00186, "system__address_imageB");
   u00187 : constant Version_32 := 16#820d6a31#;
   pragma Export (C, u00187, "system__address_imageS");
   u00188 : constant Version_32 := 16#0936cab5#;
   pragma Export (C, u00188, "system__memoryB");
   u00189 : constant Version_32 := 16#e96a4b1e#;
   pragma Export (C, u00189, "system__memoryS");

   --  BEGIN ELABORATION ORDER
   --  ada%s
   --  interfaces%s
   --  system%s
   --  system.arith_64%s
   --  system.case_util%s
   --  system.case_util%b
   --  system.htable%s
   --  system.img_enum_new%s
   --  system.img_enum_new%b
   --  system.img_int%s
   --  system.img_int%b
   --  system.img_lli%s
   --  system.img_lli%b
   --  system.io%s
   --  system.io%b
   --  system.linux%s
   --  system.multiprocessors%s
   --  system.os_primitives%s
   --  system.os_primitives%b
   --  system.parameters%s
   --  system.parameters%b
   --  system.crtl%s
   --  interfaces.c_streams%s
   --  interfaces.c_streams%b
   --  system.restrictions%s
   --  system.restrictions%b
   --  system.standard_library%s
   --  system.exceptions%s
   --  system.exceptions%b
   --  system.storage_elements%s
   --  system.storage_elements%b
   --  system.stack_checking%s
   --  system.stack_checking%b
   --  system.stack_checking.operations%s
   --  system.stack_usage%s
   --  system.string_hash%s
   --  system.string_hash%b
   --  system.htable%b
   --  system.strings%s
   --  system.strings%b
   --  system.traceback_entries%s
   --  system.traceback_entries%b
   --  ada.exceptions%s
   --  system.arith_64%b
   --  ada.exceptions.is_null_occurrence%s
   --  ada.exceptions.is_null_occurrence%b
   --  system.soft_links%s
   --  system.stack_checking.operations%b
   --  system.traces%s
   --  system.traces%b
   --  system.unsigned_types%s
   --  system.img_biu%s
   --  system.img_biu%b
   --  system.img_llb%s
   --  system.img_llb%b
   --  system.img_llu%s
   --  system.img_llu%b
   --  system.img_llw%s
   --  system.img_llw%b
   --  system.img_uns%s
   --  system.img_uns%b
   --  system.img_wiu%s
   --  system.img_wiu%b
   --  system.val_int%s
   --  system.val_lli%s
   --  system.val_llu%s
   --  system.val_uns%s
   --  system.val_util%s
   --  system.val_util%b
   --  system.val_uns%b
   --  system.val_llu%b
   --  system.val_lli%b
   --  system.val_int%b
   --  system.wch_con%s
   --  system.wch_con%b
   --  system.wch_cnv%s
   --  system.wch_jis%s
   --  system.wch_jis%b
   --  system.wch_cnv%b
   --  system.wch_stw%s
   --  system.wch_stw%b
   --  ada.exceptions.last_chance_handler%s
   --  ada.exceptions.last_chance_handler%b
   --  system.address_image%s
   --  system.bit_ops%s
   --  system.bit_ops%b
   --  system.concat_2%s
   --  system.concat_2%b
   --  system.concat_3%s
   --  system.concat_3%b
   --  system.concat_4%s
   --  system.concat_4%b
   --  system.concat_5%s
   --  system.concat_5%b
   --  system.concat_6%s
   --  system.concat_6%b
   --  system.exception_table%s
   --  system.exception_table%b
   --  ada.io_exceptions%s
   --  ada.numerics%s
   --  ada.tags%s
   --  ada.streams%s
   --  interfaces.c%s
   --  system.multiprocessors%b
   --  interfaces.c.strings%s
   --  system.crtl.runtime%s
   --  system.os_interface%s
   --  system.os_interface%b
   --  system.interrupt_management%s
   --  system.stream_attributes%s
   --  system.stream_attributes%b
   --  system.task_info%s
   --  system.task_info%b
   --  system.task_primitives%s
   --  system.interrupt_management%b
   --  system.tasking%s
   --  system.task_primitives.operations%s
   --  system.tasking%b
   --  system.tasking.debug%s
   --  system.tasking.debug%b
   --  system.task_primitives.operations%b
   --  system.traces.tasking%s
   --  system.traces.tasking%b
   --  ada.calendar%s
   --  ada.calendar%b
   --  ada.calendar.delays%s
   --  ada.calendar.delays%b
   --  system.memory%s
   --  system.memory%b
   --  system.standard_library%b
   --  system.random_numbers%s
   --  ada.numerics.float_random%s
   --  ada.numerics.float_random%b
   --  system.secondary_stack%s
   --  interfaces.c.strings%b
   --  interfaces.c%b
   --  ada.tags%b
   --  system.soft_links%b
   --  system.stack_usage%b
   --  system.secondary_stack%b
   --  system.random_numbers%b
   --  system.address_image%b
   --  system.finalization_root%s
   --  system.finalization_root%b
   --  system.finalization_implementation%s
   --  system.finalization_implementation%b
   --  ada.finalization%s
   --  ada.finalization%b
   --  ada.finalization.list_controller%s
   --  ada.finalization.list_controller%b
   --  system.file_control_block%s
   --  system.file_io%s
   --  system.os_lib%s
   --  system.os_lib%b
   --  system.file_io%b
   --  system.soft_links.tasking%s
   --  system.soft_links.tasking%b
   --  system.tasking.entry_calls%s
   --  system.tasking.initialization%s
   --  system.tasking.initialization%b
   --  system.tasking.protected_objects%s
   --  system.tasking.protected_objects%b
   --  system.tasking.utilities%s
   --  system.traceback%s
   --  ada.exceptions%b
   --  system.traceback%b
   --  ada.real_time%s
   --  ada.real_time%b
   --  ada.text_io%s
   --  ada.text_io%b
   --  ada.text_io.generic_aux%s
   --  ada.text_io.generic_aux%b
   --  ada.text_io.integer_aux%s
   --  ada.text_io.integer_aux%b
   --  ada.integer_text_io%s
   --  ada.integer_text_io%b
   --  system.tasking.protected_objects.entries%s
   --  system.tasking.protected_objects.entries%b
   --  system.tasking.queuing%s
   --  system.tasking.queuing%b
   --  system.tasking.utilities%b
   --  system.tasking.rendezvous%s
   --  system.tasking.protected_objects.operations%s
   --  system.tasking.protected_objects.operations%b
   --  system.tasking.rendezvous%b
   --  system.tasking.entry_calls%b
   --  system.tasking.stages%s
   --  system.tasking.stages%b
   --  exercise7%b
   --  END ELABORATION ORDER

end ada_main;
