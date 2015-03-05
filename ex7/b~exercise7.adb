pragma Ada_95;
pragma Source_File_Name (ada_main, Spec_File_Name => "b~exercise7.ads");
pragma Source_File_Name (ada_main, Body_File_Name => "b~exercise7.adb");

with System.Restrictions;

package body ada_main is
   pragma Warnings (Off);

   procedure Do_Finalize;
   pragma Import (C, Do_Finalize, "system__standard_library__adafinal");

   Local_Priority_Specific_Dispatching : constant String := "";
   Local_Interrupt_States : constant String := "";

   procedure adainit is
      E159 : Boolean; pragma Import (Ada, E159, "system__stack_usage_E");
      E015 : Boolean; pragma Import (Ada, E015, "system__soft_links_E");
      E025 : Boolean; pragma Import (Ada, E025, "system__exception_table_E");
      E080 : Boolean; pragma Import (Ada, E080, "ada__io_exceptions_E");
      E113 : Boolean; pragma Import (Ada, E113, "ada__numerics_E");
      E055 : Boolean; pragma Import (Ada, E055, "ada__tags_E");
      E064 : Boolean; pragma Import (Ada, E064, "ada__streams_E");
      E082 : Boolean; pragma Import (Ada, E082, "interfaces__c_E");
      E084 : Boolean; pragma Import (Ada, E084, "interfaces__c__strings_E");
      E141 : Boolean; pragma Import (Ada, E141, "system__task_info_E");
      E008 : Boolean; pragma Import (Ada, E008, "ada__calendar_E");
      E006 : Boolean; pragma Import (Ada, E006, "ada__calendar__delays_E");
      E117 : Boolean; pragma Import (Ada, E117, "system__random_numbers_E");
      E019 : Boolean; pragma Import (Ada, E019, "system__secondary_stack_E");
      E073 : Boolean; pragma Import (Ada, E073, "system__finalization_root_E");
      E075 : Boolean; pragma Import (Ada, E075, "system__finalization_implementation_E");
      E071 : Boolean; pragma Import (Ada, E071, "ada__finalization_E");
      E092 : Boolean; pragma Import (Ada, E092, "ada__finalization__list_controller_E");
      E090 : Boolean; pragma Import (Ada, E090, "system__file_control_block_E");
      E069 : Boolean; pragma Import (Ada, E069, "system__file_io_E");
      E087 : Boolean; pragma Import (Ada, E087, "system__os_lib_E");
      E171 : Boolean; pragma Import (Ada, E171, "system__tasking__initialization_E");
      E163 : Boolean; pragma Import (Ada, E163, "system__tasking__protected_objects_E");
      E121 : Boolean; pragma Import (Ada, E121, "ada__real_time_E");
      E063 : Boolean; pragma Import (Ada, E063, "ada__text_io_E");
      E169 : Boolean; pragma Import (Ada, E169, "system__tasking__protected_objects__entries_E");
      E177 : Boolean; pragma Import (Ada, E177, "system__tasking__queuing_E");
      E185 : Boolean; pragma Import (Ada, E185, "system__tasking__stages_E");

      Main_Priority : Integer;
      pragma Import (C, Main_Priority, "__gl_main_priority");
      Time_Slice_Value : Integer;
      pragma Import (C, Time_Slice_Value, "__gl_time_slice_val");
      WC_Encoding : Character;
      pragma Import (C, WC_Encoding, "__gl_wc_encoding");
      Locking_Policy : Character;
      pragma Import (C, Locking_Policy, "__gl_locking_policy");
      Queuing_Policy : Character;
      pragma Import (C, Queuing_Policy, "__gl_queuing_policy");
      Task_Dispatching_Policy : Character;
      pragma Import (C, Task_Dispatching_Policy, "__gl_task_dispatching_policy");
      Priority_Specific_Dispatching : System.Address;
      pragma Import (C, Priority_Specific_Dispatching, "__gl_priority_specific_dispatching");
      Num_Specific_Dispatching : Integer;
      pragma Import (C, Num_Specific_Dispatching, "__gl_num_specific_dispatching");
      Main_CPU : Integer;
      pragma Import (C, Main_CPU, "__gl_main_cpu");
      Interrupt_States : System.Address;
      pragma Import (C, Interrupt_States, "__gl_interrupt_states");
      Num_Interrupt_States : Integer;
      pragma Import (C, Num_Interrupt_States, "__gl_num_interrupt_states");
      Unreserve_All_Interrupts : Integer;
      pragma Import (C, Unreserve_All_Interrupts, "__gl_unreserve_all_interrupts");
      Zero_Cost_Exceptions : Integer;
      pragma Import (C, Zero_Cost_Exceptions, "__gl_zero_cost_exceptions");
      Detect_Blocking : Integer;
      pragma Import (C, Detect_Blocking, "__gl_detect_blocking");
      Default_Stack_Size : Integer;
      pragma Import (C, Default_Stack_Size, "__gl_default_stack_size");
      Leap_Seconds_Support : Integer;
      pragma Import (C, Leap_Seconds_Support, "__gl_leap_seconds_support");

      procedure Install_Handler;
      pragma Import (C, Install_Handler, "__gnat_install_handler");

      Handler_Installed : Integer;
      pragma Import (C, Handler_Installed, "__gnat_handler_installed");
   begin
      Main_Priority := -1;
      Time_Slice_Value := -1;
      WC_Encoding := 'b';
      Locking_Policy := ' ';
      Queuing_Policy := ' ';
      Task_Dispatching_Policy := ' ';
      System.Restrictions.Run_Time_Restrictions :=
        (Set =>
          (False, False, False, False, False, False, False, False, 
           False, False, False, False, False, False, False, False, 
           False, False, False, False, False, False, False, False, 
           False, False, False, False, False, False, False, False, 
           False, False, False, False, False, False, False, False, 
           False, False, False, False, False, False, False, False, 
           False, False, False, False, False, False, False, False, 
           False, False, False, False, False, False, False, True, 
           False, False, False, False, False, False, False, False, 
           False),
         Value => (0, 0, 0, 0, 0, 0, 0),
         Violated =>
          (True, False, True, True, False, False, False, True, 
           True, True, True, True, True, False, False, True, 
           False, False, True, True, False, True, True, True, 
           True, True, True, True, True, True, False, True, 
           False, True, True, False, True, True, False, True, 
           False, True, False, False, True, False, False, False, 
           False, True, False, True, False, True, True, True, 
           False, False, True, False, True, True, True, True, 
           True, False, True, False, False, True, False, True, 
           False),
         Count => (1, 0, 0, 1, 0, 1, 0),
         Unknown => (False, False, False, True, False, True, False));
      Priority_Specific_Dispatching :=
        Local_Priority_Specific_Dispatching'Address;
      Num_Specific_Dispatching := 0;
      Main_CPU := -1;
      Interrupt_States := Local_Interrupt_States'Address;
      Num_Interrupt_States := 0;
      Unreserve_All_Interrupts := 0;
      Zero_Cost_Exceptions := 1;
      Detect_Blocking := 0;
      Default_Stack_Size := -1;
      Leap_Seconds_Support := 0;

      if Handler_Installed = 0 then
         Install_Handler;
      end if;

      System.Stack_Usage'Elab_Spec;
      System.Exception_Table'Elab_Body;
      E025 := True;
      Ada.Io_Exceptions'Elab_Spec;
      E080 := True;
      Ada.Numerics'Elab_Spec;
      E113 := True;
      Ada.Tags'Elab_Spec;
      Ada.Streams'Elab_Spec;
      E064 := True;
      Interfaces.C'Elab_Spec;
      Interfaces.C.Strings'Elab_Spec;
      System.Task_Info'Elab_Spec;
      E141 := True;
      Ada.Calendar'Elab_Spec;
      Ada.Calendar'Elab_Body;
      E008 := True;
      Ada.Calendar.Delays'Elab_Body;
      E006 := True;
      E084 := True;
      E082 := True;
      Ada.Tags'Elab_Body;
      E055 := True;
      System.Soft_Links'Elab_Body;
      E015 := True;
      E159 := True;
      System.Secondary_Stack'Elab_Body;
      E019 := True;
      System.Random_Numbers'Elab_Body;
      E117 := True;
      System.Finalization_Root'Elab_Spec;
      E073 := True;
      System.Finalization_Implementation'Elab_Spec;
      System.Finalization_Implementation'Elab_Body;
      E075 := True;
      Ada.Finalization'Elab_Spec;
      E071 := True;
      Ada.Finalization.List_Controller'Elab_Spec;
      E092 := True;
      System.File_Control_Block'Elab_Spec;
      E090 := True;
      System.Os_Lib'Elab_Body;
      E087 := True;
      System.File_Io'Elab_Body;
      E069 := True;
      System.Tasking.Initialization'Elab_Body;
      E171 := True;
      System.Tasking.Protected_Objects'Elab_Body;
      E163 := True;
      Ada.Real_Time'Elab_Spec;
      Ada.Real_Time'Elab_Body;
      E121 := True;
      Ada.Text_Io'Elab_Spec;
      Ada.Text_Io'Elab_Body;
      E063 := True;
      System.Tasking.Protected_Objects.Entries'Elab_Spec;
      E169 := True;
      System.Tasking.Queuing'Elab_Body;
      E177 := True;
      System.Tasking.Stages'Elab_Body;
      E185 := True;
   end adainit;

   procedure adafinal is
   begin
      Do_Finalize;
   end adafinal;

   function main
     (argc : Integer;
      argv : System.Address;
      envp : System.Address)
      return Integer
   is
      procedure initialize (Addr : System.Address);
      pragma Import (C, initialize, "__gnat_initialize");

      procedure finalize;
      pragma Import (C, finalize, "__gnat_finalize");

      procedure Ada_Main_Program;
      pragma Import (Ada, Ada_Main_Program, "_ada_exercise7");

      SEH : aliased array (1 .. 2) of Integer;

      Ensure_Reference : aliased System.Address := Ada_Main_Program_Name'Address;
      pragma Volatile (Ensure_Reference);

   begin
      gnat_argc := argc;
      gnat_argv := argv;
      gnat_envp := envp;

      Initialize (SEH'Address);
      adainit;
      Break_Start;
      Ada_Main_Program;
      Do_Finalize;
      Finalize;
      return (gnat_exit_status);
   end;

--  BEGIN Object file/option list
   --   ./exercise7.o
   --   -L./
   --   -L/usr/lib/gcc/x86_64-linux-gnu/4.6/adalib/
   --   -shared
   --   -lgnarl-4.6
   --   -lgnat-4.6
   --   -lpthread
--  END Object file/option list   

end ada_main;
