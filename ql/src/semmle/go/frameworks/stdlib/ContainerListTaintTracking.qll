/**
 * Provides classes modeling security-relevant aspects of the standard libraries.
 */

import go

/** Provides models of commonly used functions in the `container/list` package. */
module ContainerListTaintTracking {
  private class ListBack extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).Back() *Element
    ListBack() { this.(Method).hasQualifiedName("container/list", "List", "Back") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ListFront extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).Front() *Element
    ListFront() { this.(Method).hasQualifiedName("container/list", "List", "Front") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ListInit extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).Init() *List
    ListInit() { this.(Method).hasQualifiedName("container/list", "List", "Init") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isReceiver() and outp.isResult()
    }
  }

  private class ListInsertAfter extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).InsertAfter(v interface{}, mark *Element) *Element
    ListInsertAfter() { this.(Method).hasQualifiedName("container/list", "List", "InsertAfter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and
      (outp.isReceiver() or outp.isResult())
    }
  }

  private class ListInsertBefore extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).InsertBefore(v interface{}, mark *Element) *Element
    ListInsertBefore() { this.(Method).hasQualifiedName("container/list", "List", "InsertBefore") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and
      (outp.isReceiver() or outp.isResult())
    }
  }

  private class ListMoveAfter extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).MoveAfter(e *Element, mark *Element)
    ListMoveAfter() { this.(Method).hasQualifiedName("container/list", "List", "MoveAfter") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ListMoveBefore extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).MoveBefore(e *Element, mark *Element)
    ListMoveBefore() { this.(Method).hasQualifiedName("container/list", "List", "MoveBefore") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ListMoveToBack extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).MoveToBack(e *Element)
    ListMoveToBack() { this.(Method).hasQualifiedName("container/list", "List", "MoveToBack") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ListMoveToFront extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).MoveToFront(e *Element)
    ListMoveToFront() { this.(Method).hasQualifiedName("container/list", "List", "MoveToFront") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ListPushBack extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).PushBack(v interface{}) *Element
    ListPushBack() { this.(Method).hasQualifiedName("container/list", "List", "PushBack") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and
      (outp.isReceiver() or outp.isResult())
    }
  }

  private class ListPushBackList extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).PushBackList(other *List)
    ListPushBackList() { this.(Method).hasQualifiedName("container/list", "List", "PushBackList") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ListPushFront extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).PushFront(v interface{}) *Element
    ListPushFront() { this.(Method).hasQualifiedName("container/list", "List", "PushFront") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and
      (outp.isReceiver() or outp.isResult())
    }
  }

  private class ListPushFrontList extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).PushFrontList(other *List)
    ListPushFrontList() {
      this.(Method).hasQualifiedName("container/list", "List", "PushFrontList")
    }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isReceiver()
    }
  }

  private class ListRemove extends TaintTracking::FunctionModel, Method {
    // signature: func (*List).Remove(e *Element) interface{}
    ListRemove() { this.(Method).hasQualifiedName("container/list", "List", "Remove") }

    override predicate hasTaintFlow(FunctionInput inp, FunctionOutput outp) {
      inp.isParameter(0) and outp.isResult()
    }
  }
}
